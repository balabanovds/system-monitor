# GRPC-web client

## Intro

gRPC is a modern, HTTP2-based protocol, that provides RPC semantics using the strongly-typed binary data format of protocol buffers across multiple languages (C++, C#, Golang, Java, Python, NodeJS, ObjectiveC, etc.)

With gRPC-Web, it is extremely easy to build well-defined, easy to reason about APIs between browser frontend code and microservices. Frontend development changes significantly:

-   no more hunting down API documentation - `.proto` is the canonical format for API contracts.
-   no more hand-crafted JSON call objects - all requests and responses are strongly typed and code-generated, with hints available in the IDE.
-   no more dealing with methods, headers, body and low level networking - everything is handled by grpc.invoke.
-   no more second-guessing the meaning of error codes - gRPC status codes are a canonical way of representing issues in APIs.
-   no more one-off server-side request handlers to avoid concurrent connections - gRPC-Web is based on HTTP2, with multiplexes multiple streams over the same connection.
-   no more problems streaming data from a server - gRPC-Web supports both 1:1 RPCs and 1:many streaming requests.
-   no more data parse errors when rolling out new binaries - backwards and forwards-compatibility of requests and responses.

In short, gRPC-Web moves the interaction between frontend code and microservices from the sphere of hand-crafted HTTP requests to well-defined user-logic methods.

## Description

This client implements two approaches in grpc-web

1. [Official approach](#official-approach) - github 4.8k stars
2. [Alternative approach](#alternative-approach) - github 3.2k stars

### <a id="official-approach"></a> Official approach

[Github](https://github.com/grpc/grpc-web)
[grpc.io](https://grpc.io/docs/languages/web/)

No matter what backend is written on, you have to use [Envoy](https://www.envoyproxy.io/) to proxy HTTP/1.1 to HTTP/2

Script to generate code [generate_classic.sh](./generate_classic.sh)

### <a id="alternative-approach"></a> Alternative approach

[Github](https://github.com/improbable-eng/grpc-web)

Backend written on Go and use grpcweb wrapper to expose HTTP/1.1 [handler](../internal/api/server.go) to grpc-web client.

Client uses:

-   `@improbable-eng/grpc-web` - a TypeScript gRPC-Web client library for browsers
-   `ts-protoc-gen` - a TypeScript plugin for the protocol buffers compiler that provides strongly typed message classes and method definitions.

Script to generate code [generate.sh](./generate.sh)

## Project setup

```
npm install
```

### Compiles and hot-reloads for development

```
yarn serve
```

### Compiles and minifies for production

```
yarn build
```
