# :construction: shell-playground :construction:

This is under heavy construction :construction: and might not ever be something at all.
But the goal is to create a GitHub page that host a playground to format and eventually run a single shell script.
And all that using GoLang wasm.

# Common commands

## Build the webassembly stuff
```
GOOS=js GOARCH=wasm go build -o assets/shell-playground.wasm ./cmd/wasm
```

## Run local dev server

> Make sure to have build the webassembly target before!
```
cd cmd/server
go run main.go
```
