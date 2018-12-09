package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

// printRing prints elements of the ring, starting with the currentMarble
func printRing(r *ring.Ring) {
	r.Do(func(p interface{}) {
		fmt.Printf("%v ", p.(int))
	})
	fmt.Println()
}

func main() {

	inputs := readInput("../real_input.txt")

	for _, scenario := range inputs[len(inputs)-1:] { // only do actually requested scenario for submitting
		// Initialize new ring with 1 0-element
		playingField := ring.New(1)
		playingField.Value = 0

		numPlayers, numMarbles := infoFromLine(scenario)
		numMarbles *= 100

		scores := make([]int, numPlayers, numPlayers)

		// fmt.Println(playingField.Len(), playingField.Value.(int))
		for marbleNumber := 1; marbleNumber <= numMarbles; marbleNumber++ {
			if marbleNumber%23 == 0 {
				currentPlayer := marbleNumber % numPlayers
				// current marble
				scores[currentPlayer] += marbleNumber
				// move back places in ring and take out element
				playingField = playingField.Move(-8)
				removedMarble := playingField.Unlink(1)
				// add element to player score
				scores[currentPlayer] += removedMarble.Value.(int)
				// currentMarble is first element following the removed one
				playingField = playingField.Next()
			} else {
				marble := &ring.Ring{Value: marbleNumber}
				playingField = playingField.Next()       // insert position is after 1 next to the currentMarble
				playingField = playingField.Link(marble) // add new element to ring
				// From godoc
				// If r and s point to different rings, linking them creates a single ring with the elements of s inserted after r.
				// The result points to the element following the last element of s after insertion.
				// => currentMarble should be the last element of s (the newly inserted marble), so move back 1 place
				playingField = playingField.Prev()
			}

			// printRing(playingField)
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
