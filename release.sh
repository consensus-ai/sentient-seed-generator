#!/bin/bash

# error output terminates this script
set -e

if [[ -z $1 || -z $2 ]]; then
  echo "Usage: $0 PRIVATE_KEY PUBLIC_KEY SEED_GENERATOR_VERSION"
  exit 1
fi

privkeyFile=$1
pubkeyFile=$2
uiVersion=${3:-v0.0.1}

# ensure we have a clean state
rm -rf bind* windows.syso output

# run the bundler
astilectron-bundler -v


# sign and verify the binaries
binaryName="sentient-seed-generator"
for os in osx linux windows; do
  (
    if [ $os = 'osx' ]; then
      appDir="output/darwin-amd64"
      binName="${binaryName}-${os}.app.zip"
      (
        cd $appDir
        mv "${binaryName}.app" "${binaryName}-${os}.app"
        zip -r $binName "${binaryName}-${os}.app"
      )
    elif [ $os = 'linux' ]; then
      appDir="output/linux-amd64"
      binName="${binaryName}-${os}"
      (
        cd $appDir
        mv "${binaryName}" "${binaryName}-${os}"
      )
    elif [ $os = 'windows' ]; then
      appDir="output/windows-amd64"
      binName="${binaryName}-${os}.exe"
      (
        cd $appDir
        mv "${binaryName}.exe" "${binaryName}-${os}.exe"
      )
    fi

    cd $appDir

    openssl dgst -sha256 -sign $privkeyFile -out $binName.sig $binName
    if [[ -n $pubkeyFile ]]; then
      openssl dgst -sha256 -verify $pubkeyFile -signature $binName.sig $binName
    fi
  )
done

echo "Done"
