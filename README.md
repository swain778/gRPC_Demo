# Go gRPC Server and Client

This is a basic gRPC server and client written in GO. 

In this I have implemented a simple gRPC server and client with the following functionality.

- simple RPC
- server-side streaming RPC
- client-side streaming RPC
- bidirectional streaming RPC

# setting up a gRPC-Go project
1. Create a new directory for your project and cd into it.

```bash
mkdir gRPC-Demo
cd gRPC-Demo
mkdir client server proto
```
2. Installing the gRPC Go plugin.

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"
```

3. Initialize a Go module.
```bash
go mod init gRPC-Demo

go mod tidy
```

4. Create the proto file with the required services and messages in the proto directory.

5. Generate .pb.go files from the proto file

depending on what path you mention in your greet.proto file, you will either run this

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

6. Create the new server and client directories and create the main.go files with necessary controller and services

# Running Specific API

### Unary
	callSayHello(client)

### Server stream
	callSayHelloServerStreaming(client, names)

### Client stream
	callSayHelloClientStreaming(client, names)

### Bidirectional stream
	callHelloBidirectionalStream(client, names)

To Run a specific API's we can comment out the rest three.

# Running the application

1. Installing the dependencies

```bash
go mod tidy
```

2. Run the server

```bash
go run server/main.go
```

3.Run the client 

```bash
go run client/main.go
```

