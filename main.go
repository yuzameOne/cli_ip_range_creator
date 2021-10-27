package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var zero, one, two, three, four, five, six int

// "static" byte array
var asciiArray = [10]byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

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

func pointIndexes(array []string) []int {

	idxSymbols := []int{*zero, *one, *two, *three, *four, *five, *six}

	for i := 0; i < len(array); i++ {

		str := array[i]
		count := 0

		for idx, vle := range str {

			if vle == 46 || vle == 45 {

				idxSymbols[count] = idx

				count++
			}
		}
	}

	return &idxSymbols
}

func readIpRangeFile(pathToFle string) []string {

	var stringRange []string

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

	_ = pointIndexes(arrayRange)

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

	// 24.10 1:33 pointIndexes возвращает локальный массив с индексами всех точек в строке и в removesPointAndDash
	// собираю строку в bytes.Buffer

	// 26.10 15:27 все фигня...... нужно с помощью make() хостануть []int на 6 элеметов  или "статический"
	// потом хостануть указатель на переменную и положить туда указатель на адрес элемета массива
	/*
			из книги от создателей языка

		Каждый компонент переменной составного типа — поле структуры или элемент
		массива — также является переменной, а значит, имеет свой адрес.
	*/
	//  делаем вот так
	// idxSymbols := make([]int,6)
	// (полная иицилизация,но будет краткая ) var p *int = &idxSymbols[i]
	// указатель *p будет хранить зачение элемена массива, которое можно будет изменять ))))

}

func main() {

	readIpRangeFile("example.txt")
	removesPointAndDash(stringRange)

}
