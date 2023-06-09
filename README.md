Video Tutorial: [Maximilien Andile's Comprehensive Guide](https://www.youtube.com/watch?v=gbrPMv_GuQY&t=722s)

# Setting Up Your Environment on Mac
To begin, ensure your development environment is properly configured by installing the necessary dependencies:

- Install the Protocol Buffers Compiler:
  ```bash
  brew install protobuf
  ```
- Install the Go Protocol Buffers plugin:
  ```bash
  brew install protoc-gen-go
  ```
- Install the gRPC plugin for Go:
  ```bash
  brew install protoc-gen-go-grpc
  ```

# Core Concepts to Keep in Mind

## Initial Setup
Start by creating a `.proto` file where you will:

- Define the messages: these represent the data structures used in gRPC communication.
- Outline the service: this should include the Remote Procedure Call (RPC) methods you intend to implement.

**Example Message Definitions:**
```proto
message Amount {
  int64 amount = 1; // Represents the smallest unit of the currency
  string currency = 2;
}

message CreateRequest {
  Amount amount = 1;
  string from = 2;
  string to = 3;
  string VATNumber = 4;
}

message CreateResponse {
  bytes pdf = 1;
  bytes docx = 2;
}
```

**Example of a gRPC Service Using the Defined Methods:**
```proto
service Invoicer {
  rpc Create(CreateRequest) returns (CreateResponse);
}
```

## Code Generation
Use `protoc`, the Protocol Buffers compiler, to generate Go code from your `.proto` file. The generated code will:

- Encode and decode your messages using Protocol Buffers.
- Handle incoming gRPC requests for your service.

**Command to Generate the Go Code:**
```bash
protoc --go_out=invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto
```

## Adapting to Changes
Any modifications to your `.proto` file necessitate recompiling your Go code. A `Makefile` can help automate this process. Please remember that manual edits to the generated Go code should be avoided, as any changes will be erased when the code is regenerated.

**Command to Execute the Makefile:**
```bash
make generate_grpc_code
```

# Tags
#grpc #golang #programming #microservices #container