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

	field := make([][]int, 301, 301)
	for i := 0; i < 301; i++ {
		field[i] = make([]int, 301, 301)
		for j := 0; j < len(field[i]); j++ {
			field[i][j] = calculateCoordinatesScore(i, j, serialNumber)
		}
	}

	maxX, maxY, size, totalPower := 0, 0, 0, 0

	// slow but steady :p

	// s is gridsize
	for s := 1; s < 300; s++ {
		// i,j are top-left coordinates
		for i := 0; i < 300-s; i++ {
			for j := 0; j < 300-s; j++ {

				score := 0
				for x := i; x < i+s; x++ {
					for y := j; y < j+s; y++ {
						score += field[x][y]
					}
				}

				if score > totalPower {
					totalPower = score
					maxX = i
					maxY = j
					size = s
				}
			}
		}

	}

	fmt.Println(maxX, maxY, size, totalPower)
}
