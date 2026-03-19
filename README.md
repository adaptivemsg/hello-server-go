# hello-server-go

## Message api

`api/hello` holds the shared message api for clients and the server.

## Run the server

```bash
go run -tags server .
```

## Use as a client

Import `hello-server-go/api/hello` without build tags.

## Beginner how-to

### Prerequisites

- Go installed (`go version`)
- Rust installed only if you want to run the Rust client (`cargo version`)

### Generate messages (optional)

If you need to regenerate Rust bindings from the Go message types:

```bash
go generate ./api/hello
```

### Run the server

```bash
go run -tags server .
```

### Try a client

Go client:

```bash
(cd ../hello-client-go && go run .)
```

Rust client:

```bash
(cd ../hello-client-rust && cargo run)
```

### Common issues

- Missing handler symbols: run with `-tags server`
- `amgen-go` not found: run `go run ../adaptivemsg-go/cmd/amgen-go` or install it
