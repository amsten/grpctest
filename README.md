# What to remember

## First steps
Create a file.proto
  - Define the messages
  - Define a service
    - This contains the RPC methods.

**Example messages:**
```proto
message Amount {
    int64 amount = 1; // Represented in the smallest unit of the currency
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

**Example gRPC service using the methods:**
```proto
service Invoicer {
    rpc Create(CreateRequest) returns (CreateResponse);
}
```

## Generate Go code
This is done using protoc
The code will encode/decode messages with Protocol Buffers.
It will handle incoming gRPC requests.

**Generating the code:**
```bash
protoc --go_out=invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto
```

## Changes
Each time the proto file is changed you need to regenerate the code.
Regeneration can be made easy using a Makefile.
Do not touch the generated code, changes will be overwritten during the next iteration.

**Running the Makefile**
```bash
make generate_grpc_code
```
