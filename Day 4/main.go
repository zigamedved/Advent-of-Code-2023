package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func getNumbersFromString(numbersString []string) []string {
	var result []string
	for i := 1; i < len(numbersString); i++ {
		if numbersString[i] != "" {
			result = append(result, numbersString[i])
		}
	}
	return result
}

func contains(nums []string, num string) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func validateCard(answer *int, line string) error {
	cardAndDraws := strings.Split(line, ":")
	numbers := strings.Split(cardAndDraws[1], "|")

	winningNumbersString := strings.Split(numbers[0], " ")
	drawnNumbersString := strings.Split(numbers[1], " ")

	winningNumbers := getNumbersFromString(winningNumbersString)
	drawnNumbers := getNumbersFromString(drawnNumbersString)

	var matches float64
	for _, el := range winningNumbers {
		if contains(drawnNumbers, el) {
			matches++
		}
	}
	if matches == 0 {
		return nil
	}

	*answer += int(math.Pow(2, matches-1))
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
		validateCard(&answer, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	fmt.Printf("Answer is: %d", answer)
}
