package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x  int
	y  int
	vx int
	vy int
}

func distance(p1 *point, p2 *point) int {
	return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
}

func totalDistance(points []*point) int {
	sum := 0
	for i := 0; i < len(points); i++ {
		for j := i; j < len(points); j++ {
			sum += distance(points[i], points[j])
		}
	}
	return sum
}

func infoFromLine(s string) *point {
	x, y, vx, vy := 0, 0, 0, 0
	cleanS := strings.Replace(s, " ", "", -1)
	fmt.Sscanf(cleanS, "position=<%d,%d>velocity=<%d,%d>", &x, &y, &vx, &vy)
	// p := point{x, y, vx, vy}
	p := point{y, x, vy, vx}
	return &p
}

// shiftPoints moves all points to a range of coordinates with 0 -> positive value
func shiftPoints(points []*point) ([]point, int, int) {
	minX, maxX, minY, maxY := 0, 0, 0, 0

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// move 0-punt to left-top
	newPoints := make([]point, 0, len(points))

	for _, p := range points {
		newP := point{
			p.x - minX,
			p.y - minY,
			p.vx,
			p.vy,
		}
		newPoints = append(newPoints, newP)
	}

	return newPoints, maxX - minX, maxY - minY
}

func doIteration(points []*point, steps int) {
	for _, p := range points {
		p.x = p.x + steps*p.vx
		p.y = p.y + steps*p.vy
	}
}

func printPoints(points []point, maxX, maxY int) {
	// fmt.Println(maxX, maxY)
	maxX += 1
	maxY += 1

	// create empty field
	field := make([][]string, maxX, maxX)
	for i := range field {
		field[i] = make([]string, maxY, maxY)
		for j := range field[i] {
			field[i][j] = "."
		}
	}

	// fmt.Println(len(field), len(field[0]))

	// add points
	for _, p := range points {
		// fmt.Println(p)
		field[p.x][p.y] = "#"
	}

	// print field
	for i := range field {
		s := ""
		for j := range field[i] {
			// fmt.Print(field[i][j])
			s += field[i][j]
		}
		fmt.Println(s)
	}

}

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

func main() {

	inputs := readInput("../real_input.txt")

	points := make([]*point, 0, len(inputs))

	for _, l := range inputs {
		p := infoFromLine(l)
		points = append(points, p)
	}
	// fmt.Println(points)

	dist := totalDistance(points)

	for index := 0; index < 20000; index++ {
		doIteration(points, 1)
		newDist := totalDistance(points)
		fmt.Println("Distance: ", dist)
		if newDist > dist { // gone too far, revert this step and stop changes
			doIteration(points, -1)
			fmt.Println("iteration", index)
			break
		}
		dist = newDist
	}
	shifted, maxX, maxY := shiftPoints(points)
	fmt.Println("Word dimenstions are ", maxX, " * ", maxY)
	_ = shifted
	_ = maxX
	_ = maxY
	printPoints(shifted, maxX, maxY)

}
