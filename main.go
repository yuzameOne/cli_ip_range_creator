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
var rangeMap = make(map[string]string, len(stringRange))

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

		stringRange = append(stringRange, line)
	}
	return stringRange
}

//TODO:

//  все это в switchRange
//  разбить строку Split вернет []string, strconv сравивать конвертить  и тут же откручивать
// 	возможно копировать по индексу (copy) концы строки

func splitRange(x func(string) []string) map[string]string {

	for i := 0; i < len(stringRange); i++ {

		str := stringRange[i]

		subSlice := strings.Split(str, "-")

		for i := 0; i < len(stringRange); i++ {

			rangeMap[subSlice[0]] = subSlice[1]
			break
		}

	}

	return rangeMap

}

func switchRange(x func(map[string]string)) {

}

func main() {

	readIpRangeFile("example.txt")

	splitRange(readIpRangeFile)
}
