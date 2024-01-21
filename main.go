package main

import (
	"log"
	"net"
	"sync"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	var mu sync.Mutex
	conns := make([]net.Conn, 0)

	push := func(msg []byte) {
		mu.Lock()
		defer mu.Unlock()

		for _, conn := range conns {
			conn.Write(msg)
		}
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("could not accept: %v", err)
			continue
		}

		userBuf := make([]byte, 32)
		conn.Write([]byte("Enter a username (max 32 chars): "))
		n, err := conn.Read(userBuf)
		if err != nil {
			log.Printf("could not read username: %v", err)
			conn.Close()
			continue
		}
		if n > 1 && userBuf[n-1] == '\n' {
			n = n - 1
		}

		mu.Lock()
		conns = append(conns, conn)
		mu.Unlock()

		go func(username string) {
			defer conn.Close()

			buf := make([]byte, 1500)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					return
				}
				push([]byte(username + ": " + string(buf[:n])))
			}
		}(string(userBuf[:n]))
	}
}
