package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	ascii table

	Dec   Char
	----------
	48		0
	49		1
	50		2
	51		3
	52		4
	53		5
	55		7
	56		8
	57		9
*/

// "static" byte array
var asciiArray = [10]byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57} //!????
var finalIpRange = make([]string, 1)
var arrayIndexCount int
var idxSymbols = make([]int, 7)
var stringRange []string

func сustomStringBuilder() {

	stringPointIndexes(stringRange)

}

func stringPointIndexes(array []string) {

	str := array[arrayIndexCount]
	count := 0

	for idx, vle := range str {

		if vle == 46 || vle == 45 {

			p := &idxSymbols[count]
			*p = idx
			count++
		}

	}
	arrayIndexCount++
}

func readIpRangeFile(pathToFle string) []string {

	f, err := os.Open(pathToFle)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		}

		line = strings.TrimRight(line, "\n")

		stringRange = append(stringRange, line)
	}

	return stringRange
}

func main() {

	readIpRangeFile("example.txt")
	сustomStringBuilder()

}
