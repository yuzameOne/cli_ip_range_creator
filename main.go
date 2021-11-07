package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var finalIpRange = make([]string, 0)
var arrayIndexCount int
var idxSymbols = make([]int, 7)
var stringRange []string

func createFinalFile() {

	file, err := os.OpenFile("finalRange.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		fmt.Printf("file not create : %s", err)
	}

	for _, vle := range finalIpRange {

		file.WriteString(vle + "\n")
	}
}

func сustomStringBuilder() {

	var cstr strings.Builder

	for i := 0; i < len(stringRange); i++ {

		stringPointIndexes(stringRange)

		str := stringRange[i]

		third := str[idxSymbols[1]+1 : idxSymbols[2]]

		t, _ := strconv.Atoi(third)

		for {

			for i := 0; i < 256; i++ {

				cstr.Reset()

				s := strconv.Itoa(i)
				nstr := strconv.Itoa(t)

				cstr.WriteString(str[:idxSymbols[1]] + "." + nstr + "." + s)

				finalIpRange = append(finalIpRange, cstr.String())

			}

			t++

			if strings.Compare(cstr.String(), str[11:]) == 0 {

				break
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
	createFinalFile()

}
