# pRPC-Go is a simple project illustrates how gPRC works with Go

## Project Setup

### Install Golang

Install Golang from [offical doc] (https://golang.org/doc/install)

### VSCode Setup

1. Install [VSCode] (https://code.visualstudio.com/)
2. Install VSCode Extensions: `vscode-proto3` and `Clang-Format`

### Install protoc 

On MacOSX, open a command line and type `brew install protobuf`

### Generate go code from proto files

The script below automatically generates go files based on messages definitions.

run `generate.sh` file

### Start Server

`go run greet/greet_server/server.go`