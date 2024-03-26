# Context and Concurrency Exploration in Go

This repository explores the usage of context and concurrency in GoLang.
<br>
The examples demonstrate practical applications of context and concurrency management.

> Make sure that you have installed go 1.22.1 or higher to run these examples.

Three main scenarios are covered:

## 1. Using `context.WithCancel`

### Booking Process

The booking process ensures that room booking proceeds only if the room is available. Otherwise, it cancels all booking processes.

To run the booking process, execute the following command:

```bash
go run booking/main.go
```

## 2. Using `http.Request.Context`

### Server Operations

The server component runs a backend server capable of processing or rejecting requests based on the request context.

To run the server, execute the following command:

```bash
go run server/main.go
```

To test this server as a client, use the following command:

```bash
curl localhost:8080
```

## 3. Using `context.WithTimeout`

### Client Request with Timeout

This example demonstrates a client application utilizing context with a timeout to enforce a request failure if the context deadline is exceeded.

To run this client example, execute the following command:

```bash
go run client/main.go
```

For a comprehensive demonstration, run this client alongside the `server/main.go`. This setup will illustrate the expected results for each side, showcasing the timeout behavior and its impact on request handling.