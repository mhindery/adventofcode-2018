package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/**************

**************/

func getCoordinates(s string) (int, int) {
	splitted := strings.Split(s, ", ")
	w, _ := strconv.Atoi(splitted[0])
	h, _ := strconv.Atoi(splitted[1])
	return w, h
}

// field [i][j][u v]
// i = width
// j = heigth
// u = identifier of nearest point
// v = distance to nearest point, unused
func createPlayingField(maxWidth int, maxHeight int) [][][]string {
	field := make([][][]string, maxWidth, maxWidth)
	for i := 0; i < maxWidth; i++ {
		field[i] = make([][]string, maxHeight, maxHeight)
		for j := 0; j < maxHeight; j++ {
			field[i][j] = make([]string, 2, 2)
			field[i][j][0] = "."
		}
	}
	return field
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

	coordinates := make(map[string][]int, 0)

	maxWidth, maxHeight := 0, 0
	coordID := 64 // use capital letters, as that's easy for now
	for _, input := range inputs {
		coordID++
		w, h := getCoordinates(input)
		coordinates[string(byte(coordID))] = []int{w, h}
		if w > maxWidth {
			maxWidth = w
		}
		if h > maxHeight {
			maxHeight = h
		}
	}

	maxHeight++
	maxWidth++

	fmt.Printf("Field size is %v w x %v h\n\n", maxWidth-1, maxHeight-1)

	field := createPlayingField(maxWidth, maxHeight)

	fmt.Println("Coordinates", coordinates)

	// printField(field)
	// fmt.Println("-------------")
	// putCoordinates(field, coordinates)
	// printField(field)
	// fmt.Println("-------------")
	fillField(field, coordinates)
	// printField(field)
	// fmt.Println("-------------")
	counts := countRegion(field)
	fmt.Println(counts)

}

func fillField(field [][][]string, coordinates map[string][]int) {
	maxAllowed := 10000

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			p := []int{i, j}
			sumDistance := 0
			for _, coord := range coordinates {
				sumDistance += manHatDist(p, coord)
			}
			if sumDistance < maxAllowed {
				field[i][j][0] = "#"
			}
		}
	}
}

func countRegion(field [][][]string) int {
	sum := 0

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j][0] == "#" {
				sum++
			}
		}
	}
	return sum
}

func putCoordinates(field [][][]string, coordinates map[string][]int) {
	for coordID, coords := range coordinates {
		w := coords[0]
		h := coords[1]
		field[w][h][0] = coordID
	}
}

func printField(field [][][]string) {
	for j := 0; j < len(field[0]); j++ {
		fmt.Print(j, " ")
		for i := 0; i < len(field); i++ {
			fmt.Print(field[i][j][0])
		}
		fmt.Println()
	}
}

func manHatDist(p []int, q []int) int {
	w := math.Abs(float64(p[0] - q[0]))
	h := math.Abs(float64(p[1] - q[1]))
	return int(w + h)
}
