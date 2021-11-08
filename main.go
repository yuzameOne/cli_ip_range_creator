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

func createFinalFile() {

	f, err := os.OpenFile("finalRange.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	for {
		s, ok := <-stringOutGorotine
		f.WriteString(s + "\n")
		if !ok {
			break
		}
	}

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

		s := <-queue

		go customStringBuilder(s)
	}
	close(queue)
}

func main() {

	readFile("example.txt")
	createFinalFile()
}
