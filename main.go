package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var stringRange []string
var rangeMap = make(map[string]string, len(stringRange))
var finalIprange []string

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

func splitRange(arrRange []string) map[string]string {

	for i := 0; i < len(arrRange); i++ {

		str := arrRange[i]

		subSlice := strings.Split(str, "-")

		rangeMap[subSlice[0]] = subSlice[1]

	}

	return rangeMap

}

//TODO:

// Строка содержит массив байтов, который, будучи созданным, является неизменяе­мым.
// Элементы байтового среза, напротив, можно свободно модифицировать

// у слайсов leftSlice и  rightSlice  конверитировть 2  и 3 элементы

// пакет bytes ,  []byte(string) bytes.Buffer


// 29.09.21   взаять string bulder  откручивать for + concatinate []  index of slice


func switchRange(rangeMap map[string]string) {

}

func main() {

	readIpRangeFile("example.txt")

	splitRange(stringRange)

	switchRange(rangeMap)
}
