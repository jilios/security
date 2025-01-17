package main

import (
	"buffio"
	"fmt"
	"log"
	"os"
)

func main(){
	// Open file and create scanner on top of it
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := buffio.NewScanner(file)

	// Default scanner is bufio.ScanLines. Let's use ScanWords.
	// Could also use a custom function of SplitFunc type.
	scanner.Split(buffio.ScanWords)

	// Scan for next token
	success := scanner.Scan()
	if success == false {
		// False on error or EOF. Check error 
		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// Get data from scan with Bytes() or Text()
	fmt.Println("Firts word found: ", scanner.Text())


	// Call scanner.Scan() manualy, or loop with for
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}