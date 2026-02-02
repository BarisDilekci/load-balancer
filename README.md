# TCP Load Balancer

This project is a simple TCP load balancer written in Go. It listens for incoming TCP connections and forwards them to a pool of backend servers using a round-robin algorithm.

## Features

- **Round-robin load balancing:** Distributes incoming connections evenly across multiple backend servers.
- **Concurrent handling:** Each client connection is proxied in a separate goroutine for high concurrency.
- **Transparent proxying:** Forwards data bidirectionally between clients and backend servers.

## How It Works

1. The load balancer listens on a specified address (default: `localhost:8080`).
2. When a client connects, the load balancer selects a backend server from the list using round-robin.
3. The connection is proxied between the client and the selected backend server.
4. Data is copied in both directions until the connection is closed.

## Configuration

You can configure the listening address and backend server addresses by modifying these variables in `main.go`:

```go
listenAddr = "localhost:8080"

server = []string{
    "localhost:5001",
    "localhost:5002",
    "localhost:5003",
}
```

## Usage

1. Start your backend servers on the specified ports (e.g., 5001, 5002, 5003).
2. Run the load balancer:

   ```sh
   go run main.go
   ```

3. Connect your clients to the load balancer (e.g., `localhost:8080`). The load balancer will forward connections to the backend servers.

## Example

You can test the load balancer using `nc` (netcat):

Start backend servers:
```sh
nc -lk 5001
nc -lk 5002
nc -lk 5003
```

Start the load balancer:
```sh
go run main.go
```

Connect a client:
```sh
nc localhost 8080
```

## License

MIT License