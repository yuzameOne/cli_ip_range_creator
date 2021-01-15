package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pathToFileUnparseIp string

func main() {

	fmt.Println("Enter path to file  unparseip :")
	fmt.Scanln(&pathToFileUnparseIp)

	file, err := os.Open(pathToFileUnparseIp)

	if err != nil {
		log.Fatalf("file don't open")
	}

	reader := bufio.NewScanner(file)

	for {

		reader.Scan()

		text := reader.Text()

		if len(text) == 0 {
			break
		}

		fmt.Println(text)

	}

}
