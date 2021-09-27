package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

// Влоить один if в  другой  if

func switchRange(rangeMap map[string]string) {

	for index, value := range rangeMap {

		compareStrings := strings.Compare(index, value)

		if compareStrings == -1 || compareStrings > 1 {

			leftSlice := strings.Split(index, ".")
			rightSlice := strings.Split(value, ".")

			if leftSlice[3] != rightSlice[3] {

				leftIndex, _ := strconv.Atoi(leftSlice[3])
				rightIndex, _ := strconv.Atoi(rightSlice[3])

				number := rightIndex - leftIndex

				for i := 0; i < number; i++ {

					leftIndex++

					s := strconv.Itoa(leftIndex)

					leftSlice[3] = s

					addArray := strings.Join(leftSlice, ".")

					finalIprange = append(finalIprange, addArray)
				}

			}

			if leftSlice[2] != rightSlice[2] {

				leftIndex, _ := strconv.Atoi(leftSlice[2])
				rightIndex, _ := strconv.Atoi(rightSlice[2])

				number := rightIndex - leftIndex

				for i := 0; i < number; i++ {

					leftIndex++

					s := strconv.Itoa(leftIndex)

					leftSlice[2] = s

					addArray := strings.Join(leftSlice, ".")

					finalIprange = append(finalIprange, addArray)
				}

			}

		}

		// fmt.Printf(" key : %s, value : %s  \n", index, value)

		// fmt.Print(value[1])
	}

}

func main() {

	readIpRangeFile("example.txt")

	splitRange(stringRange)

	switchRange(rangeMap)
}
