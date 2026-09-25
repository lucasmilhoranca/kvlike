# kvlike

A Redis-like in-memory key-value store written in Go. It serves a plain-text protocol over TCP and supports key expiration, optional append-only file (AOF) persistence, and a small interactive client. The codebase follows a clean architecture layout with dependency injection via Google Wire.

## Requirements

- Go 1.26 or later
- Docker and Docker Compose (optional)

## Running locally

Clone the repository and download dependencies:

```sh
git clone https://github.com/lucasmilhoranca/kvlike.git
cd kvlike
go mod download
```

Start the server:

```sh
go run ./cmd/server
```

Available flags:

| Flag    | Default | Description                                       |
| ------- | ------- | ------------------------------------------------- |
| `-port` | `6379`  | Port to listen on                                 |
| `-aof`  | `false` | Enable AOF persistence (writes to `data.aof`)     |

Example with persistence on a custom port:

```sh
go run ./cmd/server -port 7000 -aof
```

In another terminal, connect with the client:

```sh
go run ./cmd/client localhost:6379
```

To build binaries instead:

```sh
go build -o server ./cmd/server
go build -o client ./cmd/client
./server -aof
```

## Running with Docker

```sh
docker compose up --build
```

The server is exposed on port `6379` with AOF persistence enabled. To open the client inside the container:

```sh
docker compose exec redis-like ./client localhost:6379
```

## Commands

| Command                   | Description                                       |
| ------------------------- | ------------------------------------------------- |
| `SET key value`           | Store a value                                     |
| `GET key`                 | Get a value, or `nil` if missing                  |
| `DEL key [key ...]`       | Delete one or more keys                           |
| `EXISTS key [key ...]`    | Count how many of the given keys exist            |
| `KEYS [pattern]`          | List keys matching a pattern                      |
| `EXPIRE key seconds`      | Set a time to live on a key                       |
| `TTL key`                 | Get the remaining time to live of a key           |
| `PERSIST key`             | Remove the expiration from a key                  |
| `PING [message]`          | Check that the server is alive                    |
| `INFO [section]`          | Show server statistics                            |
| `QUIT`                    | Close the connection                              |

Example session:

```
SET name kvlike
OK
EXPIRE name 60
OK
TTL name
60
GET name
kvlike
```

## Project structure

```
cmd/
  server/           TCP server entrypoint
  client/           Interactive client
internal/
  domain/           Entities, commands and repository interfaces
  usecase/          Command handling and statistics
  adapter/          TCP handler and protocol parser
  infraestructure/  In-memory storage and AOF persistence
  container/        Wire dependency injection setup
```
