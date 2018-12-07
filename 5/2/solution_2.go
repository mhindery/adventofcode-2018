package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

/**************

**************/

func cutLettersString(s string) string {
	for {
		for index := 1; index < len(s); index++ {
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

	asciA, asciZ := 65, 91

	minLength := len(input)
	minCutset := ""

	var wg sync.WaitGroup
	var mut sync.Mutex

	for asciiCode := asciA; asciiCode < asciZ; asciiCode++ {
		wg.Add(1)

		go func(asciiCode int) {
			defer wg.Done()
			cutSet := string(byte(asciiCode)) + string(byte(asciiCode+32))

			cleaned := strings.Replace(input, string(byte(asciiCode)), "", -1)
			cleaned = strings.Replace(cleaned, string(byte(asciiCode+32)), "", -1)

			res := cutLettersString(cleaned)

			fmt.Println(cutSet, len(res))

			mut.Lock()
			if len(res) < minLength {
				minLength = len(res)
				minCutset = cutSet
			}
			mut.Unlock()

		}(asciiCode)
	}

	wg.Wait()

	fmt.Println("=> Cutset", minCutset, "results in length", minLength)
}
