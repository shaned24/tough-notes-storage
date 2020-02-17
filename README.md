# Tough Notes Storage

 A notes storage application using [gRPC](#grpc) and [protocol buffers](#protocol-buffers) with a MongoDB backend


## Contents
- [Tough Notes Storage](#tough-notes-storage)
  - [Contents](#contents)
    - [Environment Variables](#environment-variables)
    - [gRPC](#grpc)
    - [Protocol Buffers](#protocol-buffers)
    - [Running the server](#running-the-server)

### Environment Variables

| NAME                     | Description                                                                                                   | **default**   |
| ------------------------ | ------------------------------------------------------------------------------------------------------------- | ------------- |
| MONGO_CONNECTION_TIMEOUT | The timeout in seconds that the mongo client should wait until it stops trying to connect to a mongo instance | **5**         |
| MONGO_HOST               | The mongo host to connect to                                                                                  | **localhost** |
| MONGO_PORT               | The port number of the mongo host to connect to                                                               | **27017**     |
| MONGO_DB_NAME            | The name of the MongoDB name to use                                                                           | **notes**     |
| MONGO_COLLECTION_NAME    | The name of the MongoDB collection to use                                                                     | **notes**     |


### gRPC

The GRPC [quick start](https://grpc.io/docs/quickstart/go/) documents the tools
needed to get started

### Protocol Buffers

The [protoc](https://github.com/protocolbuffers/protobuf) compiler is required to generate gRPC code.

### Running the server

```bash
go run .
```