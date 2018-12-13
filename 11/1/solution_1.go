package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(loc string) []string {
	inputs := make([]string, 0)

	file, err := os.Open(loc)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputs
}

func calculateCoordinatesScore(x int, y int, serialNumber int) int {
	rackID := x + 10
	score := (rackID * y) + serialNumber
	score = score * rackID
	score = score % 1000
	score = (score - (score % 100)) / 100
	return score - 5
}

func main() {

	// inputs := readInput("../real_input.txt")
	serialNumber := 1308

	field := make([][]int, 300, 300)
	for i := 0; i < 300; i++ {
		field[i] = make([]int, 300, 300)
		for j := 0; j < len(field[i]); j++ {
			field[i][j] = calculateCoordinatesScore(i, j, serialNumber)
		}
	}

	maxX, maxY, totalPower := 0, 0, 0

	for i := 0; i < 300-3; i++ {
		for j := 0; j < len(field[i])-3; j++ {
			score := field[i][j] + field[i][j+1] + field[i][j+2]
			score += field[i+1][j] + field[i+1][j+1] + field[i+1][j+2]
			score += field[i+2][j] + field[i+2][j+1] + field[i+2][j+2]
			if score > totalPower {
				totalPower = score
				maxX = i
				maxY = j
			}
		}
	}

	fmt.Println(maxX, maxY, totalPower)

	// res := calculateCoordinatesScore(3, 5, serialNumber)
	// fmt.Println(res)

}
