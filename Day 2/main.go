package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const redLimit = 12
const greenLimit = 13
const blueLimit = 14

func validateGame(answer *int, line string, index int) error {
	gameAndDraws := strings.Split(line, ":")
	draws := strings.Split(gameAndDraws[1], ";")

	for _, dr := range draws {
		for _, nac := range strings.Split(dr, ",") { // nac=number and color: " 5 red"
			s := strings.Split(nac, " ") // split at " ", get 3 parts, second part is digit, third is color
			num, err := strconv.Atoi(s[1])
			if err != nil {
				return fmt.Errorf("Conversion error:", err)
			}

			switch s[2] {
			case "red":
				if num > redLimit {
					return nil
				}
				continue
			case "green":
				if num > greenLimit {
					return nil
				}
				continue
			case "blue":
				if num > blueLimit {
					return nil
				}
				continue
			default:
				return fmt.Errorf("Unknown color!")
			}
		}
	}

	*answer += index
	return nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	answer := 0
	gameIndex := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameIndex++
		validateGame(&answer, line, gameIndex)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	fmt.Printf("Answer is: %d", answer)
}
