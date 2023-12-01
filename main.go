package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func calculateLineSum(answer *int, line string) error {

	var start, end int = 0, len(line)
	var first, last rune

	if start == end {
		return nil
	}

	for _, ch := range line {
		if unicode.IsDigit(ch) {
			first = ch
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		ch := []rune(line)[i]
		if unicode.IsDigit(ch) {
			last = ch
			break
		}
	}

	sum, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		fmt.Println("Conversion error:", err)
		return err
	}

	*answer += sum
	return nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	answer := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		calculateLineSum(&answer, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	fmt.Printf("Sum of first and last digit is: %d", answer)
}
