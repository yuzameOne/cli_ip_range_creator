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

func splitRange(arrRange []string) map[string]string {

	for i := 0; i < len(arrRange); i++ {

		str := arrRange[i]

		strings.Trim(str, "\n")

		subSlice := strings.Split(str, "-")

		subSlice[1] = strings.TrimRight(subSlice[1], "\n")

		rangeMap[subSlice[0]] = subSlice[1]

	}

	return rangeMap

}

//TODO:

//  все это в switchRange
//  разбить строку Split вернет []string, strconv сравивать конвертить  и тут же откручивать
// 	возможно копировать по индексу (copy) концы строки

func switchRange(rangeMap map[string]string) {

	for index, value := range rangeMap {

		fmt.Printf(" key : %s, value : %s  \n", index, value)

		fmt.Print(value[1])
	}

}

func main() {

	readIpRangeFile("example.txt")

	splitRange(stringRange)

	switchRange(rangeMap)
}
