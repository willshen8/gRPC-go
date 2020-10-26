# gRPC-Go 

This a simple project illustrates how gPRC works with Go with server and client implementations.

## Project Structure

```
root
| - greet/
    | -- greet_client
    | -- greet_server
    | -- greetpb
| - calculator/
    | -- calculator_client
    | -- calculator_server
    | -- calculatorpb

```
client folder contains code that a server sends requests from.
server folder contains code that a server sends responds to requests.

## Prerequisite

There are few things you'll need to setup to get going:

### Install Golang

Install Golang from [offical doc](https://golang.org/doc/install)

### VSCode Setup

1. Install [VSCode](https://code.visualstudio.com/)
2. Install VSCode Extensions: `vscode-proto3` and `Clang-Format`

### Install protoc

On MacOSX, open a command line and type `brew install protobuf`

### Install gRPC

`go get -u google.golang.org/grpc`

## Generate go code from proto files

The script below automatically generates go proto file based on messages definitions.

run `generate.sh` file

### Start Server
e.g. 
`go run greet/greet_server/server.go`

### Start Client
e.g. 
`go run greet/greet_client/client.go`