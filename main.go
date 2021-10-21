package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var stringRange []string
var newCustomString strings.Builder

// var subSlice []string
// var pureSlice []string

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

func removesPointAndDash(arrayRange []string) {

	str := "5.100.67.0-5.100.67.255"

	point := bytes.Index([]byte(str), []byte("-"))

	leftString := str[:point]
	rightString := str[point+1:]

	fmt.Println(leftString)
	fmt.Println(rightString)

	// TODO
	//  что надо  67.0 и  67.255 из стиринги в  инты сравить и обратно в строку  bytes.Buffer
	
	// 21.10 14:11 стоял на смене и пришла мысль в голову , что последнее число  всегда будет в диапапозоне
	//  от 0 до 255 (67.0). нужно просто создать Enum от 0 до 255 и подставить в цикле в строку  bytes.Buffer
}

//  TODO
//  писать subSlice в  [][]string

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
	removesPointAndDash(stringRange)

}
