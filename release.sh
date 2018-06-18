#!/bin/bash

# error output terminates this script
set -e

if [[ -z $1 || -z $2 ]]; then
  echo "Usage: $0 PRIVATE_KEY PUBLIC_KEY SEED_GENERATOR_VERSION"
  exit 1
fi

privkeyFile=$1
pubkeyFile=$2
version=${3:-v0.0.1}

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

  binarySuffix="${version}-${os}-amd64"
  compiledBinaryName=$compiledBinaryBaseName$binaryExtension
  outputBinaryName="${compiledBinaryBaseName}-${binarySuffix}$binaryExtension"

  mv $appDir/$compiledBinaryName $outputDir/$outputBinaryName

  (
    cd $outputDir
    zipFile=$outputBinaryName.zip
    zip -r $zipFile $outputBinaryName

    openssl dgst -sha256 -sign $privkeyFile -out $zipFile.sig $zipFile
    if [[ -n $pubkeyFile ]]; then
      openssl dgst -sha256 -verify $pubkeyFile -signature $zipFile.sig $zipFile
    fi
  )
done

echo "Done"
