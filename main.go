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

// "static" byte array
var asciiArray = [10]byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

/*
	Dec   Char
	----------
	48		0
	49		1
	50		2
	51		3
	52		4
	53		5
	54		6
	55		7
	56		8
	57		9
*/

func pointIndexes(str string) {

	str := "5.100.67.0-5.100.67.255"

	for idx, vle := range str {

		if vle == 46 || vle == 45{

		}
	}

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

	// 22.10 3:06  81 это (56 49) а 82 это (56 50). 112 (49 49 50) - 123(49 50 51).
	//  циферки в таблице ascii от 0 до 9  в таблице ascii  от 48  до 57
	//  ТО ЕСТЬ  увеличиваем крайний байт до  0 (112 - 119  потом  сбрасываем на 0(48) 110) с увеличением предыдущего
	// байта на +1 (110 [49 49 48] до 120 [49 50 48])

	// 22.10 14:18 написать маленькую функицию которая будет проходить по строке (for index := range str) и посчитает все
	// точки и тире. в итоге будет 7 индексов, 3-тий индекс будет тире
	//  не понимаю как вернуть зачение!!!!!! писать в массив, а потом на каждой итерации удалять все элементы(дорого) или
	// пихать в переменую (кортежи) результать фукнкции (7 идексов)

	// 23.10 13:54 буду пихать в кортеж, потому-что  не происходит выделение памяти при переопределении
	// функция pointIndexes хостануть локальный "статический" массив и положить туда сслыки на переменные  вот так
	//  [&one,&two,&three .... &ten]  и count откручивать индекс и return кортеж

	/*


		one, two := "0", "1"

		fmt.Printf("one : %s, two : %s  \n", one, two)

		fmt.Println(&one, &two)

		one, two = "2", "3"

		fmt.Printf("one : %s, two : %s  \n", one, two)

		fmt.Println(&one, &two)


			вывод в консоли.
			one : 0, two : 1
			0xc00010a040 0xc00010a050
			one : 2, two : 3
			0xc00010a040 0xc00010a050


	*/
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
