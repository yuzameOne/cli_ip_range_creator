package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var finalIpRange = make([]string, 1)
var arrayIndexCount int
var idxSymbols = make([]int, 7)
var stringRange []string

// first string(stringRange) in idxSymbols  [1 5 8 10 12 16 19] 5.100.67.0-5.100.67.255

func сustomStringBuilder() {

	var cstr strings.Builder

	for i := 0; i < len(stringRange); i++ {

		stringPointIndexes(stringRange)

		str := stringRange[i]

		third := str[idxSymbols[1]+1 : idxSymbols[2]]

		seventh := str[idxSymbols[5]+1 : idxSymbols[6]]

		t, _ := strconv.Atoi(third)
		s, _ := strconv.Atoi(seventh)

		d := s - t

		if d == 0 {

			for i := 0; i < 256; i++ {

				s := strconv.Itoa(i)
				cstr.WriteString(str[:9] + s)

				finalIpRange = append(finalIpRange, cstr.String())

				cstr.Reset()
			}

		} else if d != 0 {

			cstr.Reset()

			for i := 0; i < d; i++ {

				s := t + i
				n := strconv.Itoa(s)
				cstr.WriteString(str[:6] + n + "." + str[9:10])

				fmt.Println(cstr.String())

			}
		}
	}
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

	fmt.Println(finalIpRange)

}
