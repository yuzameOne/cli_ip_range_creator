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
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		}

		stringRange = append(stringRange, line)
	}
	return stringRange
}

//TODO:
// не поимаю как ложить  слайс в  многомерный слайс
// index 2,3 and 6,7

func clearExtraChar(x func(string) []string) {

	for i := 0; i < len(stringRange); i++ {

		str := stringRange[i]

		for index, value := range str {

			if value == 46 || value == 45 {

				subSliceStringBuilder = append(subSliceStringBuilder, newCustomString.String())
				newCustomString.Reset()
			}
			if value != 46 && value != 45 {

				newCustomString.WriteByte(str[index])

				if len(str)-1 == index {
					subSliceStringBuilder = append(subSliceStringBuilder, newCustomString.String())
					newCustomString.Reset()
				}

			}
		}
		sliceOfSlice = append(sliceOfSlice, subSliceStringBuilder)
	}
	fmt.Println(sliceOfSlice)
}

func main() {

	readIpRangeFile("example.txt")

	clearExtraChar(readIpRangeFile)
}
