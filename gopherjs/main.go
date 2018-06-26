package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash"

	mnemonics "github.com/NebulousLabs/entropy-mnemonics"
	"github.com/NebulousLabs/fastrand"
	"github.com/NebulousLabs/merkletree"
	"github.com/gopherjs/gopherjs/js"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ed25519"
)

// NewSenHash creates a new instance of SenHash with blake512
// as the underlying hasher.
func NewSenHash() hash.Hash {
	b2b, _ := blake2b.New512(nil)
	return SenHash{
		blake512hasher: b2b,
	}
}

// SenHash is an abstraction to make blake-512 hasher return a 256-bit prefix.
// This is really a pretty lazy solution driven by lack of proper time and resources
// to fully switch the entire codebase to using 512-bit hashes and entropy.
type SenHash struct {
	blake512hasher hash.Hash
}

// Write adds data to the underlying hasher.
// Part of the io.Writer interface
func (h SenHash) Write(p []byte) (n int, err error) {
	return h.blake512hasher.Write(p)
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (h SenHash) Sum(b []byte) []byte {
	b = append(b, h.blake512hasher.Sum(nil)[:32]...)
	return b
}

// Reset resets the Hash to its initial state.
func (h SenHash) Reset() {
	h.blake512hasher.Reset()
}

// Size returns the number of bytes Sum will return.
func (h SenHash) Size() int {
	return 32
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (h SenHash) BlockSize() int {
	return h.BlockSize()
}

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

type SeedGenerator struct {
	seed    string
	address string
}

func (s *SeedGenerator) Seed() string {
	return s.seed
}

func (s *SeedGenerator) Address() string {
	return s.address
}

func (s *SeedGenerator) Generate() {
	wallet, err := GenerateWallet()
	if err != nil {
		return
	}

	s.seed = wallet.SeedPhrase
	s.address = wallet.FirstPubAddr
}

func NewSeedGenerator() *js.Object {
	seedGenerator := SeedGenerator{}
	seedGenerator.Generate()

	return js.MakeWrapper(&seedGenerator)
}

func main() {
	js.Global.Set("seedgenerator", map[string]interface{}{
		"NewSeedGenerator": NewSeedGenerator,
	})
}
