package main

import (
	"net"
	"testing"
	"time"
)

func TestTelnet(t *testing.T) {
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Cannot start test server: %v", err)
	}
	defer server.Close()

	go func() {
		conn, _ := server.Accept()
		defer conn.Close()
		conn.Write([]byte("Anybody here?"))
	}()

	timeout := 1 * time.Second
	host := "localhost"
	port := "8080"

	dialer := &net.Dialer{
		Timeout: timeout,
	}

	conn, err := dialer.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		t.Fatalf("Cannot connect: %v", err)
	}
	defer conn.Close()

	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("Error occurred while receiving data: %v", err)
	}

	expected := "Anybody here?"
	if string(buf[:n]) != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, string(buf[:n]))
	}
}
