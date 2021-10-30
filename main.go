package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

// first string(stringRange) in idxSymbols  [1 5 8 10 12 16 19] 5.100.67.0-5.100.67.255

func сustomStringBuilder() {
	// var buf bytes.Buffer

	// for i := 0; i < len(stringRange); i++ {

	stringPointIndexes(stringRange)

	// str := stringRange[i]

	var cstr strings.Builder

	str := "5.100.75.0-5.100.76.255"

	third := str[idx[1]+1 : idx[2]]

	seventh := str[idx[5]+1 : idx[6]]

	t, _ := strconv.Atoi(third)
	s, _ := strconv.Atoi(seventh)

	d := -(t - s)

	if str[idx[1]+1:idx[2]] != str[idx[5]+1:idx[6]] {

		for i := 0; i < 256; i++ {

			cstr.Reset()
			s := strconv.Itoa(i)
			cstr.WriteString(str[:9] + s)

			fmt.Println(cstr.String())
		}

	}

	fmt.Printf("%v  \n", cstr)

	fmt.Printf("%v \n", d)
	// fmt.Println(cstr.String())

}

// }

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

	fmt.Println(idxSymbols)
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
