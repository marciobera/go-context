# Context and Concurrency Exploration in Go

This repository explores the usage of context and concurrency in GoLang.
<br>
The examples demonstrate practical applications of context and concurrency management.

Two main scenarios are covered:

## 1. Using `context.WithCancel`

### Booking Process

The booking process ensures that room booking proceeds only if the room is available. Otherwise, it cancels all booking processes.

To run the booking process, execute the following command:

```bash
go run booking/main.go
```

## 2. Utilizing `http.Request.Context`

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