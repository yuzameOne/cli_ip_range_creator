package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

var THREADS int = runtime.GOMAXPROCS(0)

/*
	TODO

	так  будет fanIn -> fanOut

	для чего нужна  переменая gomaxprocs ???? для того чтобы задейвствовать все потоки которые есть у планировщика на проце автоматически

	в моем случае это 4 потока(ноут), дома настольный ПК  16 потоков.

	


*/

func readFile(fname string) {

	f, err := os.Open(fname)

	if err != nil {

		fmt.Println(err)
	}

	rf := bufio.NewScanner(f)

	for rf.Scan() {

		fmt.Println(rf.Text())
	}
}

func main() {

	readFile("example.txt")

}
