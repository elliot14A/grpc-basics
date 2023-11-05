# gRPC Server and Client in Go

This is a simple gRPC project in Go that demonstrates a gRPC server running on port 8080 and a gRPC client, which is an HTTP server running on port 8081. The project follows the Ports and Adapters design pattern for a clean separation of concerns.

## Project Structure

The project structure is organized as follows:

```
totality-corp-assessment/
    ├── main.go         # Entry point for running the server and client
    ├── server/         # GRPC Server-side code
    ├── client/         # GRPC Client-side code and http rest server code
    └── proto/          # Protocol Buffers (protobuf) definitions
```

## How to Run

### Start the gRPC Server:

```bash
go run main.go server
```

This command will start the gRPC server on port 8080.

### Start the gRPC Client (HTTP Server):

```bash
go run main.go client
```

This command will start the gRPC client, which is an HTTP server, on port 8081.

**Note: All the envs used in the project can be viewed inside config.env**

## Available Routes

The gRPC client exposes the following routes:

- `/users`: Get a list of users.
- `/users/{id}`: Get user details by ID.
