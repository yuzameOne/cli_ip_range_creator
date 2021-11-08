package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var threads int = runtime.GOMAXPROCS(0)
var queue = make(chan string, threads)
var stringOutGorotine = make(chan string, threads)

/*
	TODO

	так  будет fanOut -> fanIn

	для чего нужна  переменая gomaxprocs ???? для того чтобы задейвствовать все потоки которые есть у планировщика на проце автоматически

	в моем случае это 4 потока(ноут), дома настольный ПК  16 потоков.

	 дальше  будет канал который будет в роли очереди (queue FIFO) , размером в количество потоков  (т.е 4)

	 канал будет принимать строку из файла и отдавать горутинам для работы


*/
func createFinalFile() {

	f, err := os.OpenFile("finalRange.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		fmt.Println(err)
	}

	f.WriteString(<- stringOutGorotine + "\n")
}

func customStringBuilder(stringOutChan string) {

	var cstr strings.Builder
	var idxSymbols = make([]int, 7)
	var count int

	for idx, vle := range stringOutChan {

		if vle == 46 || vle == 45 {

			p := &idxSymbols[count]
			*p = idx
			count++
		}

	}

	third := stringOutChan[idxSymbols[1]+1 : idxSymbols[2]]

	t, _ := strconv.Atoi(third)

	for {

		for i := 0; i < 256; i++ {

			cstr.Reset()

			s := strconv.Itoa(i)
			nstr := strconv.Itoa(t)

			cstr.WriteString(stringOutChan[:idxSymbols[1]] + "." + nstr + "." + s)

			stringOutGorotine <- cstr.String()
		}

		t++

		if strings.Compare(cstr.String(), stringOutChan[11:]) == 0 {

			break
		}

	}
}

func readFile(fname string) {

	f, err := os.Open(fname)

	defer f.Close()

	if err != nil {

		fmt.Println(err)
	}

	rf := bufio.NewScanner(f)

	for rf.Scan() {

		queue <- rf.Text()
	}
}

func main() {

	readFile("example.txt")

	for s := range queue {

		go customStringBuilder(s)
	}

}
