package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Open a new file for writing only
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	byteSlice := []byte("Bytes:22\n")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes. \n", byteWritten)

	// Variadic second
	err = ioutil.WriteFile("test_copy.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
