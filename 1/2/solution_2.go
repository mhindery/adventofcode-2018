package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	reachedFrequencies := make(map[int]int)
	frequency := 0
	index := -1
	for {
		index = (index + 1) % len(inputs)
		num, _ := strconv.Atoi(inputs[index])
		frequency += num
		if _, ok := reachedFrequencies[frequency]; ok {
			fmt.Println("Reached frequency ", frequency, "again")
			break
		}
		reachedFrequencies[frequency]++
	}
}
