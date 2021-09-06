package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var newCustomString strings.Builder
var subSlice []string
var stringRange []string

func main() {
	f, err := os.Open("example.txt")

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

	fmt.Println(stringRange)

	// str := "5.100.67.0-5.100.67.255"

	// for index, value := range str {

	// 	if value == 46 || value == 45 {

	// 		subSlice = append(subSlice, newCustomString.String())
	// 		newCustomString.Reset()
	// 	}

	// 	if value != 46 && value != 45 {

	// 		newCustomString.WriteByte(str[index])

	// 		if len(str)-1 == index {
	// 			subSlice = append(subSlice, newCustomString.String())
	// 		}

	// 	}

	// }
	// fmt.Println(subSlice)

	//TODO:

	// index 2,3 and 6,7
	// нужно в for сравнивать СТРОКИ index 2,3 and 6,7 если меньше  конвертировать в int, инкрементить
	// и тут же конверитить в стороку и собирать это в нову строку StringBulder

	//  СУКА 2 фукции 3 канала))))) и горутины)))
}