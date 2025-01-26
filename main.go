package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	counter int

	// TODO configurable
	listenAddr = "localhost:8080"

	// TODO configurable
	server = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal("failed to listen: %s ", err)
	}
	defer listener.Close()
	fmt.Printf("TCP Proxy server is now listening on port: %s", listenAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept: %s ", err)
		}

		backend := chooseBackend()
		fmt.Println("backend: %s", backend)
		go func() {
			err := proxy(backend, conn)
			log.Printf("Warning: proxying failed: %v", err)
		}()
	}
}

func proxy(backend string, conn net.Conn) error {
	bc, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed to connect to backend %s: %v", backend, err)
	}
	go io.Copy(conn, bc)

	go io.Copy(bc, conn)

	return nil
}

func chooseBackend() string {
	s := server[counter%len(server)]
	counter++
	return s
}
