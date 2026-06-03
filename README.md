# Client–Server (Go)

A minimal TCP chat demo in Go: one server accepts multiple clients, and each client sends line-based messages from the terminal.

## Requirements

- [Go](https://go.dev/dl/) 1.18 or newer

## Project layout

```
clientserver/
├── server/server.go   # TCP listener, handles clients concurrently
└── client/client.go   # Interactive TCP client
```

## Usage

### 1. Start the server

Pass the listen address (host:port or `:port`):

```bash
cd server
go run server.go :8080
```

### 2. Start a client

In another terminal, pass the server IP and port:

```bash
cd client
go run client.go 127.0.0.1 8080
```

You can open several clients; the server handles each connection in its own goroutine.

## Behavior

- The server sends `Welcome` when a client connects.
- Type a message and press Enter; the server logs it and replies with `Received: '<your message>'`.
- Type `exit` (case-insensitive) on the client to close the connection gracefully.

## Example

**Server**

```
Server listening on port:  :8080
--- New Connection ---
[Message from 127.0.0.1:54321]: hello
```

**Client**

```
Type a message to send: hello
[127.0.0.1:8080 -> 127.0.0.1:54321] Server confirmation: Received: 'hello'
```

