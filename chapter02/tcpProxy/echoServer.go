package main

// Note: The following program is refusing connections
// Possibly because of windows firewall?

import (
	"io"
	"log"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close()
	// creating a buffer to hold the transmitted data.
	buf := make([]byte, 1024)

	for {
		size, err := conn.Read(buf[0:])
		if err == io.EOF {
			log.Println("[+]\tClient Disconnected. Exiting.")
			break
		}
		if err != nil {
			log.Fatalln("[!]\tUnexpected Error. Disconnecting...")
			break
		}
		// Printing info about the recv data
		log.Println("[~]\tReceived %d bytes.: %s", size, string(buf))

		// Echo back
		log.Println("[+]\tEchoing...")
		_, err = conn.Write(buf)
		if err != nil {
			log.Fatalln("[!]\tUnable to write data. Exiting...")
			break
		}
	}
}

// Listen, then accept connection then send to handler.
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatalln("[!]\tCould not bind to port 20000. Exiting...")
		return
	}
	log.Printf("[+]\tListener established on 0.0.0.0:20000")
	for {
		conn, err := listener.Accept()
		log.Println("[+]\tReceived connection.")
		if err != nil {
			log.Fatalln("[!]\tUnable to accept connection.")
		}
		go handler(conn)
	}
}
