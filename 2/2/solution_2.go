package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processPair(s string, t string) (found bool) {
	common := ""

	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			common += string(s[i])
		}
	}

	if len(s)-len(common) == 1 {
		fmt.Println(common)
		found = true
	}
	return
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputs := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs); j++ {
			if processPair(inputs[i], inputs[j]) {
				os.Exit(0)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
