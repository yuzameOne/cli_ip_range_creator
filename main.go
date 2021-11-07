package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

var threads int = runtime.GOMAXPROCS(0)

var queue = make(chan string, threads)

/*
	TODO

	так  будет fanIn -> fanOut

	для чего нужна  переменая gomaxprocs ???? для того чтобы задейвствовать все потоки которые есть у планировщика на проце автоматически

	в моем случае это 4 потока(ноут), дома настольный ПК  16 потоков.

	 дальше  будет канал который будет в роли очереди (queue FIFO) , размером в количество потоков  (т.е 4)

	 канал будет принимать строку из файла и отдавать горутинам для работы


*/

func readFile(fname string) {

	f, err := os.Open(fname)

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

}
