package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processLine(s string) (countTwo int, countThree int) {
	runeCount := make(map[rune]int)

	for _, letter := range s {
		runeCount[letter]++
	}

	for _, count := range runeCount {
		if count == 2 {
			countTwo = 1
		} else if count == 3 {
			countThree = 1
		}
	}

	return
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	countTwo := 0
	countThree := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		two, three := processLine(scanner.Text())
		countTwo += two
		countThree += three
	}

	fmt.Println(countThree * countTwo)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
