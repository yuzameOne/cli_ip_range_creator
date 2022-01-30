package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var threads int = runtime.NumCPU()
var queue = make(chan string, threads)
var stringOutGorotine = make(chan string, threads)

var wg sync.WaitGroup

func createFinalFile(fname string) {

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

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

	if stringOutChan[:idxSymbols[3]] == stringOutChan[idxSymbols[3]+1:] || stringOutChan[idxSymbols[1]+1:idxSymbols[2]] > stringOutChan[idxSymbols[5]+1:idxSymbols[6]] {
		wg.Done()

		return
	}

	for {

		for i := 0; i < 256; i++ {

			cstr.Reset()

			s := strconv.Itoa(i)
			nstr := strconv.Itoa(t)

			cstr.WriteString(stringOutChan[:idxSymbols[1]] + "." + nstr + "." + s)

			stringOutGorotine <- cstr.String()
		}

		t++

		if strings.Compare(cstr.String(), stringOutChan[idxSymbols[3]+1:idxSymbols[6]]+".255") == 0 {

			break
		}

	}
	wg.Done()
}

func readFile(fname string) {

	f, err := os.Open(fname)

	defer f.Close()

	if err != nil {

		fmt.Println(err)
	}

	rf := bufio.NewScanner(f)

	for rf.Scan() {

		queue <- strings.TrimRight(rf.Text(), "\n")

	}

	close(queue)
}

func main() {

	start := time.Now()

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	var argOne string
	flag.StringVar(&argOne, "ptf", "", "path to file ")

	if argOne != "" {
		fmt.Println("the first argument is missing : path to file ")
		os.Exit(3)
	}

	var argTwo string
	flag.StringVar(&argTwo, "ptsf", "", "path to  save file ")

	flag.Parse()

	if argTwo == "" {
		argTwo = path + "/new_" + argOne
	}

	fmt.Println("Start work")

	go readFile(argOne)

	go createFinalFile(argTwo)

	for {

		s, ok := <-queue

		if !ok {
			break
		}

		wg.Add(1)
		go customStringBuilder(s)

	}

	wg.Wait()

	duration := time.Since(start)

	time.Sleep(duration)

	close(stringOutGorotine)

	fmt.Printf("Time to work : %v , file name : new_%s  \n", duration, argOne)

}
