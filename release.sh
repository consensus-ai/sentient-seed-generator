#!/bin/bash

# error output terminates this script
set -e

if [[ -z $1 || -z $2 ]]; then
  echo "Usage: $0 PRIVATE_KEY PUBLIC_KEY SEED_GENERATOR_VERSION [CODE_SIGN_IDENTITY]"
  exit 1
fi

privkeyFile=$1
pubkeyFile=$2
version=${3:-v0.0.1}
codesignIdent=$4

# ensure we have a clean state
rm -rf bind* windows.syso output

# run the bundler
echo "Bundling the electron app..."
astilectron-bundler
echo "Done bundling the electron app"

# sign and verify the binaries
compiledBinaryBaseName="sentient-seed-generator"
outputDir="output"

for os in osx linux windows; do
  if [ $os = 'osx' ]; then
    appDir="output/darwin-amd64"
    binaryExtension=".app"
  elif [ $os = 'linux' ]; then
    appDir="output/linux-amd64"
    binaryExtension=""
  elif [ $os = 'windows' ]; then
    appDir="output/windows-amd64"
    binaryExtension=".exe"
  fi

  outputBinaryName=$compiledBinaryBaseName$binaryExtension

  mv $appDir/$outputBinaryName $outputDir/$outputBinaryName

  (
    cd $outputDir
    zipFile=$compiledBinaryBaseName-${version}-${os}-amd64.zip

    # OSX is very 'special'
    if [ $os = 'osx' ]; then
      # First we ned to open the app - otherwise the signatures produced below will be broken.
      # This is because go-astilectron-bundler does some sketchy stuff on first boot, and expands
      # the bundle, which invalidates the signatures.
      # echo "Open the binary and close it after it has loaded. Because of reasons."
      open $outputBinaryName &
      sleep 3
      killall sentient-seed-generator

      cp ../bundle_resources/osx/* $outputBinaryName/Contents
      rm -rf $outputBinaryName/Contents/MacOS/vendor/*.zip

      if [[ -n $codesignIdent ]]; then
        # sign the app on OSX
        # NOTE: you can create a self-signed cert here via keychain access > certificate assistant > create a certificate
        echo "signing OSX bundle"
        codesign --deep --force --sign $codesignIdent $outputBinaryName

        echo "verifying OSX bundle signature"
        # verify signature
        codesign --verify -v $outputBinaryName
        spctl -a -vvvv $outputBinaryName
      else
        echo "Code signing identity not specified. Creating unsigned binary."
      fi

      # use special zip util on OSX to preserve the file metadata with its signature
      ditto -c -k --sequesterRsrc --keepParent $outputBinaryName $zipFile
    else
      zip -ry $zipFile $outputBinaryName
    fi

    openssl dgst -sha256 -sign $privkeyFile -out $zipFile.sig $zipFile
    if [[ -n $pubkeyFile ]]; then
      openssl dgst -sha256 -verify $pubkeyFile -signature $zipFile.sig $zipFile
    fi
  )
done

echo "Done"
