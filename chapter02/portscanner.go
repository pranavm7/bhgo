package main

import (
	"fmt"
	"net"
)

func main() {
	for port := 1; port <= 1024; port++ {
		go func(port int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				panic(err)
			}
			conn.Close()
			fmt.Printf("%d open\n", port)
		}(port)
	}

}
