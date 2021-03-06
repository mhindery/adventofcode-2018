package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/**********
waaaaaaay to slow, after an hour and looked to do better than naive slice of ints ...
**********/

func infoFromLine(s string) (int, int) {
	splitted := strings.Split(s, " ")
	numPlayers, _ := strconv.Atoi(splitted[0])
	numMarbles, _ := strconv.Atoi(splitted[6])
	return numPlayers, numMarbles
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

func getLoopedIndex(i int, arr []int) int {
	ret := i
	if ret < 0 {
		ret += len(arr)
	}
	if ret >= len(arr) {
		ret = ret % (len(arr))
	}
	// fmt.Println("Called with", i, "/", arr, "(", len(arr), ")", "=>", ret)
	return ret
}

func main() {

	inputs := readInput("../real_input.txt")

	for _, scenario := range inputs {

		numPlayers, numMarbles := infoFromLine(scenario)
		numMarbles *= 100

		scores := make([]int, numPlayers, numPlayers)

		marbles := make([]int, numMarbles+1, numMarbles+1)
		for index := 0; index < numMarbles+1; index++ {
			marbles[index] = index
		}
		// marbles[len(marbles)-1] += 100
		// fmt.Println(marbles[len(marbles)-1])
		// marbles[len(marbles)-1] += 100
		// fmt.Println(marbles[len(marbles)-1])

		playingField := make([]int, 0)
		currentPlayer := 0
		currentMarbleIndex := 0

		for len(marbles) > 0 {
			newMarble := marbles[0]
			// fmt.Println("Processing", newMarble)

			if newMarble != 0 && newMarble%23 == 0 {
				scores[currentPlayer] += newMarble
				toRemoveMarbleIndex := getLoopedIndex(currentMarbleIndex-7, playingField)
				removedValue := playingField[toRemoveMarbleIndex]
				scores[currentPlayer] += removedValue
				playingField = append(playingField[:toRemoveMarbleIndex], playingField[toRemoveMarbleIndex+1:]...)
				currentMarbleIndex = toRemoveMarbleIndex
				// fmt.Println("Kept", newMarble, "removed", removedValue, "at", toRemoveMarbleIndex, ":", playingField, "cmi =", currentMarbleIndex, "/", playingField[currentMarbleIndex])
			} else if len(playingField) == 0 {
				playingField = append(playingField, newMarble)
			} else {
				insertIndex := getLoopedIndex(currentMarbleIndex+1, playingField) + 1
				if insertIndex > len(playingField) {
					insertIndex = insertIndex % len(playingField)
				}

				playingField = append(playingField[:insertIndex], append([]int{newMarble}, playingField[insertIndex:]...)...)
				// fmt.Println("Inserting", newMarble, "at", insertIndex, ":", playingField, "cmi =", insertIndex)
				currentMarbleIndex = insertIndex
			}

			marbles = marbles[1:]
			currentPlayer = (currentPlayer + 1) % len(scores)
		}

		score := 0
		for _, val := range scores {
			if val > score {
				score = val
			}
		}
		fmt.Println("Scenario: ", scenario, ": score", score)
	}

}
