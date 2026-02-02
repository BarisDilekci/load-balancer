package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync/atomic"
)

var (
	counter uint64

	listenAddr = "localhost:8080"

	server = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatal("failed to listen %s", err)

	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		backend := chooseBackend()
		log.Printf("new connection -> %s", backend)

		go func(c net.Conn, b string) {
			if err := proxy(b, c); err != nil {
				log.Printf("proxy error (%s): %v", b, err)
			}
		}(conn, backend)
	}

}

func proxy(backend string, client net.Conn) error {
	backendConn, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("connect backend %s failed: %w", backend, err)
	}

	defer client.Close()
	defer backendConn.Close()

	errCh := make(chan error, 2)

	go func() {
		_, err := io.Copy(backendConn, client)
		errCh <- err
	}()

	go func() {
		_, err := io.Copy(client, backendConn)
		errCh <- err
	}()

	err = <-errCh
	return err
}

func chooseBackend() string {
	i := atomic.AddUint64(&counter, 1)
	return server[i%uint64(len(server))]
}
