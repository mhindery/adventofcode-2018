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
	counts := countPlanes(field)
	fmt.Println("Plane counts", counts)
	counts = filterPlanes(field, counts)
	fmt.Println("Filtered plane counts", counts)

	id, max := maxCount(counts)
	fmt.Println("Max plane size", max, "with id", id)
}

func maxCount(counts map[string]int) (string, int) {
	max := 0
	id := "."
	for k, v := range counts {
		if v > max {
			max = v
			id = k
		}
	}
	return id, max
}

func fillField(field [][][]string, coordinates map[string][]int) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			coord := []int{i, j}
			coordID, distance := getNearestPoint(coord, coordinates)
			_ = distance
			field[i][j][0] = coordID
		}
	}
}

// countPlanes counts the size of each plane
func countPlanes(field [][][]string) map[string]int {
	counts := make(map[string]int)
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			coordID := field[i][j][0]
			counts[coordID]++
		}
	}
	return counts
}

// filterPlanes removes a plane from the counts if it is at the border
func filterPlanes(field [][][]string, counts map[string]int) map[string]int {
	maxWidth := len(field) - 1
	maxHeigth := len(field[0]) - 1

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			// Discard coordinates for planes which extend infinitely
			if i == 0 || i == maxWidth || j == 0 || j == maxHeigth {
				coordID := field[i][j][0]
				if _, ok := counts[coordID]; ok {
					delete(counts, coordID)
				}
			}
		}
	}
	return counts
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

// getNearestPoint gets the id and distance to the nearest coordinate from the given set
// if 2 points are equally near, the id will be empty
func getNearestPoint(p []int, coordinates map[string][]int) (string, int) {
	maxDistance := 999999999
	coordID := "."

	for id, coord := range coordinates {
		distance := manHatDist(p, coord)
		if distance < maxDistance {
			coordID = id
			maxDistance = distance
		} else if distance == maxDistance {
			coordID = "."
		}
	}
	return coordID, maxDistance
}

func manHatDist(p []int, q []int) int {
	w := math.Abs(float64(p[0] - q[0]))
	h := math.Abs(float64(p[1] - q[1]))
	return int(w + h)
}
