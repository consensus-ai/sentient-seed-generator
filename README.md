# Development Setup

## 1. Install the repo and dependencies
```bash
go get -u github.com/consensus-ai/sentient-seed-generator/...
go get -u github.com/asticode/go-astilectron/...
go get -u github.com/asticode/go-astilectron-bootstrap/...
go get -u github.com/asticode/go-astilectron-bundler/...
```

## 2. Bundle the app for your environment
```bash
cd $GOPATH/src/github.com/consensus-ai/sentient-seed-generator
astilectron-bundler -v
```

## 3. Open the app
Open the bundle for your OS located in the `output` directory.
