package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var newCustomString strings.Builder
var subSliceStringBuilder []string
var sliceOfSlice [][]string
var stringRange []string

func readIpRangeFile(pathToFle string) []string {

	f, err := os.Open(pathToFle)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString(10)

		if err == io.EOF {
			break
		}

		stringRange = append(stringRange, line)
	}
	return stringRange
}

//TODO:

// идея с могомерным массивом не удалась (((((
// переписать название  фукции  clearExtraChar

// что думаю : разбить строку Split вернет []string, strconv сравивать конвертить  и тут же откручивать

func clearExtraChar(x func(string) []string) {

}

func main() {

	readIpRangeFile("example.txt")

	clearExtraChar(readIpRangeFile)
}
