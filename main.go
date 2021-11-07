package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(fname string) {

	f, err := os.Open(fname)

	if err != nil {

		fmt.Println(err)
	}

	rf := bufio.NewScanner(f)

	for rf.Scan() {

		fmt.Println(rf.Text())
	}
}

func main() {

	readFile("example.txt")

}
