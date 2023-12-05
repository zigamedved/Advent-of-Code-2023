package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getNeighbors(mp [][]rune, start, finish, rowIndex int) []rune { // get's neighbors
	var result []rune
	var lookAbove, lookBelow bool

	if rowIndex > 0 {
		lookAbove = true
	}
	if rowIndex < len(mp)-1 {
		lookBelow = true
	}

	for i := start; i <= finish; i++ {
		if lookAbove {
			result = append(result, mp[rowIndex-1][i])
		}
		if lookBelow {
			result = append(result, mp[rowIndex+1][i])
		}
	}

	if lookAbove && start > 0 { // left up
		result = append(result, mp[rowIndex-1][start-1])
	}
	if lookAbove && finish < len(mp[0])-1 { // right up
		result = append(result, mp[rowIndex-1][finish+1])
	}
	if lookBelow && start > 0 { // left down
		result = append(result, mp[rowIndex+1][start-1])
	}
	if lookBelow && finish < len(mp[0])-1 { // right down
		result = append(result, mp[rowIndex+1][finish+1])
	}

	if start > 0 {
		result = append(result, mp[rowIndex][start-1])
	}

	if finish < len(mp[0])-1 {
		result = append(result, mp[rowIndex][finish+1])
	}

	return result
}

func containsSymbol(chars []rune) bool {
	for _, ch := range chars {
		if !unicode.IsDigit(ch) && ch != 46 {
			return true
		}

	}
	return false
}

func readFileIntoMatrix(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	matrix := [][]rune{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

func main() {
	filePath := "input.txt"
	matrix, err := readFileIntoMatrix(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	answer := 0
	for rowIndex, row := range matrix {
		var number string
		var start, end int = -1, -1
		for ix, el := range row {
			var skip bool = false
			if ix == len(row)-1 && unicode.IsDigit(row[ix]) {
				skip = true
				number += string(el)
				ix++
			}
			if unicode.IsDigit(el) && skip == false {
				if start == -1 {
					start = ix
				}
				number += string(el)
			} else if number != "" {
				end = ix - 1
				//fmt.Printf("%s is a number, starts at %d and ends at %d in %d row \n", number, start, end, rowIndex)
				neighbors := getNeighbors(matrix, start, end, rowIndex)
				if containsSymbol(neighbors) {
					num, err := strconv.Atoi(number)
					if err != nil {
						os.Exit(1)
					}
					answer += num
				}

				number = ""
				start = -1
				end = -1
			}
		}
	}

	fmt.Println("Answer is:", answer)
}
