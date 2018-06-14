package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/NebulousLabs/entropy-mnemonics"
	"github.com/NebulousLabs/fastrand"
	"github.com/NebulousLabs/merkletree"
	"golang.org/x/crypto/ed25519"
)

type senPublicKey struct {
	Algorithm [16]byte
	Key       []byte
}

type unlockConditions struct {
	Timelock           uint64
	PublicKeys         []senPublicKey
	SignaturesRequired uint64
}

type Wallet struct {
	SeedPhrase   string
	FirstPubAddr string
}

func GenerateWallet() (Wallet, error) {
	seed := generateSeed()

	seedPhrase, err := seedToString(seed)
	if err != nil {
		return Wallet{}, err
	}

	firstPubAddr := seedToFirstPubAddr(seed)

	return Wallet{
		SeedPhrase:   seedPhrase,
		FirstPubAddr: firstPubAddr,
	}, nil
}

func generateSeed() [32]byte {
	var seed [32]byte
	fastrand.Read(seed[:])
	return seed
}

// SeedToString converts a wallet seed to a human friendly string.
func seedToString(seed [32]byte) (string, error) {
	hasher := NewSenHash()
	hasher.Write(seed[:])
	fullChecksum := hasher.Sum(nil)
	checksumedSeed := append(seed[:], fullChecksum[:6]...)

	phrase, err := mnemonics.ToPhrase(checksumedSeed, "english")
	if err != nil {
		return "", fmt.Errorf("Failed to convert seed to phrase. Reason: %+v\n", err)
	}

	return phrase.String(), nil
}

func seedToFirstPubAddr(seed [32]byte) string {
	hasher := NewSenHash()
	hasher.Write(seed[:])
	hasher.Write(encodeInt64(0))
	firstHash := hasher.Sum(nil)

	pk, _, _ := ed25519.GenerateKey(bytes.NewReader(firstHash))
	senPk := senPublicKey{
		Algorithm: [16]byte{'e', 'd', '2', '5', '5', '1', '9'},
		Key:       pk,
	}
	unlockConditions := unlockConditions{
		PublicKeys:         []senPublicKey{senPk},
		SignaturesRequired: 1,
	}
	unlockHash := unlockHash(unlockConditions)

	// get the checksum
	checksumHasher := NewSenHash()
	checksumHasher.Write(unlockHash)
	fullChecksum := checksumHasher.Sum(nil)

	return fmt.Sprintf("%x%x", unlockHash[:], fullChecksum[:6])
}

// EncInt64 encodes an int64 as a slice of 8 bytes.
func encodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

// EncInt64 encodes a uint64 as a slice of 8 bytes.
func encodeUint64(u uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, u)
	return b
}

func unlockHash(uc unlockConditions) []byte {
	hasher := NewSenHash()
	tree := merkletree.New(hasher)
	tree.Push(encodeUint64(uc.Timelock))

	var spkBuf bytes.Buffer
	spkBuf.Write(uc.PublicKeys[0].Algorithm[:])

	var lenBuf [8]byte
	binary.LittleEndian.PutUint64(lenBuf[:], uint64(len(uc.PublicKeys[0].Key)))
	spkBuf.Write(lenBuf[:])

	spkBuf.Write(uc.PublicKeys[0].Key)

	tree.Push(spkBuf.Bytes())
	tree.Push(encodeUint64(uc.SignaturesRequired))

	return tree.Root()
}
