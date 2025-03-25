package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	// Обработка параметров командной строки
	var inputFile, outputFile string
	var workers int
	flag.StringVar(&inputFile, "i", "", "Input file path (required)")
	flag.StringVar(&outputFile, "o", "new_output.txt", "Output file path")
	flag.IntVar(&workers, "w", 4, "Number of workers")
	flag.Parse()

	if inputFile == "" {
		fmt.Println("Input file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Проверка существования входного файла
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("Input file %s does not exist\n", inputFile)
		os.Exit(1)
	}

	// Открытие файлов с обработкой ошибок
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	// Каналы для параллельной обработки
	lines := make(chan string, workers*2)
	results := make(chan string, workers*2)
	var wg sync.WaitGroup

	// Пул worker'ов
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range lines {
				if processed, ok := processLine(line); ok {
					results <- processed
				}
			}
		}()
	}

	// Горутина для записи результатов
	go func() {
		writer := bufio.NewWriter(outFile)
		defer writer.Flush()
		for res := range results {
			writer.WriteString(res + "\n")
		}
	}()

	// Чтение входного файла
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		lines <- strings.TrimSpace(scanner.Text())
	}
	close(lines)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}

	// Ожидание завершения обработки
	wg.Wait()
	close(results)

	fmt.Printf("Processing completed in %v\n", time.Since(start))
}

func processLine(line string) (string, bool) {
	parts := strings.FieldsFunc(line, func(r rune) bool { 
		return r == '.' || r == '-' 
	})

	// Проверка корректности входных данных
	if len(parts) < 6 {
		return "", false
	}

	// Проверка на идентичность диапазонов
	if strings.Join(parts[:3], ".") == strings.Join(parts[3:6], ".") {
		return "", false
	}

	startIP, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", false
	}

	endIP, err := strconv.Atoi(parts[4])
	if err != nil {
		return "", false
	}

	// Генерация результатов
	var result strings.Builder
	for i := startIP; i <= endIP; i++ {
		for j := 0; j < 256; j++ {
			result.WriteString(fmt.Sprintf("%s.%d.%d\n", parts[0], i, j))
		}
	}

	return result.String(), true
}
