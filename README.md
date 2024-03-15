# :construction: shell-playground :construction:

This is under heavy construction :construction: and might not ever be something at all.
But the goal is to create a GitHub page that host a playground to format and eventually run a single shell script.
And all that using GoLang wasm.

# Common commands

## Build the webassembly stuff
```
pushd cmd/wasm
GOOS=js GOARCH=wasm -o ../../assets/
popd
```

## Run local dev server

>> Make sure to have build the webassembly target before!
```
go run cmd/server/main.go
```
