package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/**************

**************/

func cutLettersString(s string) string {
	for {
		for index := 1; index < len(s); index++ {
			// fmt.Println(index)
			doCut := (strings.ToUpper(string(s[index])) == strings.ToUpper(string(s[index-1]))) && !(s[index] == s[index-1])
			if doCut {
				// fmt.Printf("Before (%v): %v\n", index, s)
				tmp := []byte(s)
				tmp = append(tmp[:index-1], tmp[index+1:]...)
				s = string(tmp)
				// fmt.Printf("After  (%v): %v\n", index, s)
				// fmt.Println()
				index = 0
			}
		}
		break
	}
	return s
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

	input := inputs[0]
	// input = input[:100]

	// fmt.Println(input)

	res := cutLettersString(input)
	fmt.Println("=== result ===")
	fmt.Println(res)
	fmt.Println(len(res))
	// 9526
}
