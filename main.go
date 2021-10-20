package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var stringRange []string
var newCustomString strings.Builder
var subSlice []string
var pureSlice []string

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

//  TODO
//  писат subSlice в  [][]string

// func removesPointAndDash(arrayRange []string) {

// 	for i := 0; i < len(arrayRange); i++ {

// 		str := arrayRange[i]

// 		newCustomString.Reset()

// 		for index, value := range str {

// 			if value == 46 || value == 45 {

// 				subSlice = append(subSlice, newCustomString.String())
// 				newCustomString.Reset()
// 			}

// 			if value != 46 && value != 45 {

// 				newCustomString.WriteByte(str[index])

// 				if len(str)-1 == index {
// 					subSlice = append(subSlice, newCustomString.String())
// 				}

// 			}

// 			fmt.Println(subSlice)

// 		}

// 		for i := 0; i < len(subSlice); {

// 			fmt.Println(subSlice)
// 			subSlice[len(subSlice)-1] = " "
// 			fmt.Println(subSlice)
// 			subSlice = subSlice[:len(subSlice)-1]
// 		}

// 		fmt.Println(subSlice)
// 	}

// }

func main() {

	readIpRangeFile("example.txt")
	// removesPointAndDash(stringRange)

}
