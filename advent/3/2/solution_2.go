package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLineInfo(s string) (x int, y int, width int, height int) {
	locInfo := strings.Split(strings.Split(s, " @ ")[1], ": ")
	x, _ = strconv.Atoi(strings.Split(locInfo[0], ",")[0])
	y, _ = strconv.Atoi(strings.Split(locInfo[0], ",")[1])
	width, _ = strconv.Atoi(strings.Split(locInfo[1], "x")[0])
	height, _ = strconv.Atoi(strings.Split(locInfo[1], "x")[1])
	return
}

func createPlayingField(inputs []string) [][]int {
	maxWidth, maxHeight := 0, 0
	for _, line := range inputs {
		x, y, width, height := getLineInfo(line)

		tileWidth := x + width
		if tileWidth > maxWidth {
			maxWidth = tileWidth
		}

		tileHeight := y + height
		if tileHeight > maxHeight {
			maxHeight = tileHeight
		}
	}

	fmt.Printf("Field is %v x %v\n", maxWidth, maxHeight)

	field := make([][]int, maxWidth, maxWidth)
	for i := 0; i < maxWidth; i++ {
		field[i] = make([]int, maxHeight, maxHeight)
	}

	return field
}

func fillPlayingField(inputs []string, field [][]int) [][]int {
	for _, line := range inputs {
		x, y, width, height := getLineInfo(line)

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				field[i][j]++
			}
		}
	}

	return field
}

func countPixelsFromCount(field [][]int, limit int) int {
	count := 0
	for _, line := range field {
		for _, value := range line {
			if value >= limit {
				count++
			}
		}
	}
	return count
}

// tryClaim returns true if the claim is not overlapping with another,
// which equals all pixels of it are value 1 (its own claim and nothing more)
func tryClaim(field [][]int, x, y, width, height int) bool {
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			if field[i][j] != 1 {
				return false
			}
		}
	}
	return true
}

func checkClaims(inputs []string, field [][]int) {
	fmt.Println("Checking")
	for _, line := range inputs {
		x, y, width, height := getLineInfo(line)
		if tryClaim(field, x, y, width, height) {
			fmt.Println(line)
			return
		}
	}
	fmt.Println("Nothing found")
}

func main() {
	inputs := make([]string, 0)

	file, err := os.Open("../real_input.txt")
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

	field := createPlayingField(inputs)
	filledField := fillPlayingField(inputs, field)
	checkClaims(inputs, filledField)
}
