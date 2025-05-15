package main

import (
	"fmt"
	"log"
	"os"
)

// Defines io.Reader
type FooReader struct{}

// Read data from stdin
func (FooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

// At large, this is a duplication of the io.Copy(&writer,&reader) func
func main() {

	var (
		Reader FooReader
		Writer FooWriter
	)

	// Creating the var to store input
	input := make([]byte, 4096)

	s, err := Reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}

	fmt.Printf("Read %d bytes from stdin.\n", s)

	// writing from the data held in the input var
	s, err = Writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout.\n", s)
}
