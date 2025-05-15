package main

import (
	"fmt"
	"net"
	"sort"
)

// iff connected to port, insert port num in results channel else 0
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100) //create 100 buffered channel
	results := make(chan int)    // create an unbuffered channel
	var openPorts []int
	// for each channel in ports run the async worker function, passing the relevant args
	for i := 0; i < cap(ports); i++ {
		// ports = size 100 buffered channel
		// results = unbuffered channel
		go worker(ports, results)
	}
	// doing this async putting values 1-1024 in ports
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	//close the channels
	close(ports)
	close(results)
	// sort for pretty printing
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("Port %d is open\n", port)
	}
}
