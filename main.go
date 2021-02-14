package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var differenceBetweenIndex []int
var convertStringSliceToInt []int
var convertIntSliceToString []string
var finishStringWriteToFile string
var parseIPLinesFile []string
var count int = 0

func main() {

	// get root path main file
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	// get and check first argument in cli
	var argOne string
	flag.StringVar(&argOne, "pathtofile", "", "path to file ")

	if argOne != "" {
		fmt.Println("the first argument is missing")
		os.Exit(3)
	}

	// get and check second argument in cli
	var argTwo string
	flag.StringVar(&argTwo, "pathtosavefile", "", "path to file ")

	flag.Parse()

	if argTwo == "" {
		argTwo = path + "/new_ip_" + argOne
	}

	//  open file
	file, err := os.Open(argOne)

	// check error
	if err != nil {
		log.Fatal("file not Open", err)
	}

	//  read file inside all bytes
	scanner := bufio.NewScanner(file)

	// read file before \n
	scanner.Split(bufio.ScanLines)

	// write newline string in []string
	for scanner.Scan() {
		parseIPLinesFile = append(parseIPLinesFile, scanner.Text())
	}

	// close file
	defer file.Close()
	defer os.Exit(3)

	//  read []string  in loop
	for i := 0; i < len(parseIPLinesFile); i++ {

		// take string like  "5.100.67.0-5.100.67.255"
		startString := parseIPLinesFile[i]
		//  split string in []string
		newString := strings.Split(startString, "-")
		// split string [] string
		subStringLeft := strings.Split(string(newString[0]), ".")
		// split string [] string
		subStringRigth := strings.Split(string(newString[1]), ".")
		// CRUTCH , take subStringLeft[2]  static value
		twoIndexStaticNumber, _ := strconv.Atoi(subStringLeft[2])

		fmt.Sprintln(subStringRigth) // compiler scold, not use variable
		fmt.Sprintln(subStringLeft)  // compiler scold, not use variable

		//  convert  subStringLeft and subStringRigth []int, equals
		//  index value and  append  in []int differenceBetweenIndex
		for i := 2; i < len(subStringLeft); i++ {

			if subStringLeft[i] != subStringRigth[i] || subStringLeft[i] == subStringRigth[i] {

				left, _ := strconv.Atoi(subStringLeft[i])
				right, _ := strconv.Atoi(subStringRigth[i])

				differenceBetweenIndex = append(differenceBetweenIndex, -(left - right))

			}

		}

		//  delete all elements slice convertStringSliceToInt
		if len(convertStringSliceToInt) != 0 {
			for i := 0; i < 4; i++ {

				convertStringSliceToInt = convertStringSliceToInt[:len(convertStringSliceToInt)-1]
			}
		}

		//  convert all string elements to int type , and append  to slice convertStringSliceToInt
		for _, val := range subStringLeft {

			value, _ := strconv.Atoi(val)

			convertStringSliceToInt = append(convertStringSliceToInt, value)

		}

		// open/create file where write finaly range ip
		file, err := os.OpenFile(argTwo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalf("file not create : %s", err)
		}
		// get three index to int slice  and increment him
		for i := convertStringSliceToInt[2]; i <= (convertStringSliceToInt[2] + differenceBetweenIndex[0]); i++ {

			if i == (twoIndexStaticNumber+differenceBetweenIndex[0])+1 {
				break
			}
			convertStringSliceToInt[2] = i

			// get get index to int slice  and increment him
			for j := convertStringSliceToInt[3]; j <= (convertStringSliceToInt[3] + differenceBetweenIndex[1]); j++ {

				convertStringSliceToInt[3] = j

				//  convert all int elements to string type , and append  to slice convertStringSliceToInt
				for _, val := range convertStringSliceToInt {

					value := strconv.Itoa(val)

					convertIntSliceToString = append(convertIntSliceToString, value)
				}
				// convert string slice to string
				finishStringWriteToFile = strings.Join(convertIntSliceToString, ".")

				// delete all element slice sitring convertIntSliceToString
				for i := 0; i < len(convertIntSliceToString); {

					convertIntSliceToString[len(convertIntSliceToString)-1] = " "
					convertIntSliceToString = convertIntSliceToString[:len(convertIntSliceToString)-1]
				}
				// write finish string to file
				file.WriteString(finishStringWriteToFile + "\n")

				// counter string
				count = count + 1

				// console output number string line in the file
				fmt.Printf("lines in file : %d \n", count)

				//  if j == 255
				if j == differenceBetweenIndex[1] {
					convertStringSliceToInt[3] = 0

					break
				}

			}
			file.Close()
		}

	}
}
