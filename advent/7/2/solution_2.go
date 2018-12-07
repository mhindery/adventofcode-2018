package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type dependancyGraph map[string][]string

// returns dependancyID, nodeID
func infoFromLine(s string) (string, string) {
	// Step B must be finished before step E can begin.
	splitted := strings.Split(s, " ")
	return splitted[1], splitted[7]
}

func readInput() []string {
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

	return inputs
}

func main() {

	inputs := readInput()

	nodeDependancies := make(dependancyGraph)

	for _, input := range inputs {
		dependancyID, nodeID := infoFromLine(input)

		if _, ok := nodeDependancies[nodeID]; !ok {
			nodeDependancies[nodeID] = make([]string, 0)
		}
		if _, ok := nodeDependancies[dependancyID]; !ok {
			nodeDependancies[dependancyID] = make([]string, 0)
		}

		nodeDependancies[nodeID] = append(nodeDependancies[nodeID], dependancyID)
		// fmt.Println(nodeID, nodeDependancies[nodeID])
	}
	printGraph(nodeDependancies)

	numWorkers := 5
	workers := make(map[string]int)

	sequence := ""
	currentTime := 0

	for len(nodeDependancies) > 0 {
		options := nodeDependancies.selectAvailableSteps()

		for _, nodeID := range options {
			if _, ok := workers[nodeID]; !ok && len(workers) < numWorkers {
				sequence += nodeID
				sbytes := []byte(nodeID)
				time := int(sbytes[0]) - 4
				workers[nodeID] = currentTime + time
				fmt.Println("starting", nodeID, "at", currentTime, "finishes at", workers[nodeID])
			}
		}

		// Skip to earliest finish time of worker
		finishTimes := []int{}
		for _, finishTime := range workers {
			finishTimes = append(finishTimes, finishTime)
		}
		sort.Ints(finishTimes)
		currentTime = finishTimes[0]

		// Remove the node which is now finished
		for nodeID := range workers {
			if workers[nodeID] == currentTime {
				nodeDependancies.removeProcessedStep(nodeID)
				delete(workers, nodeID)
				fmt.Println("finished", nodeID, "at", currentTime)
			}
		}

		// time.Sleep(2 * time.Second)
	}

	fmt.Println("Took", currentTime, "iterations:", sequence)

}

func printGraph(nodeDependancies dependancyGraph) {
	fmt.Println("---------------")
	for nodeID, dependancies := range nodeDependancies {
		intval := []byte(nodeID)
		intval[0] -= 4
		fmt.Printf("%v\t%v\t%v\n", intval[0], nodeID, dependancies)
	}
	fmt.Println("---------------")
}

func (dg dependancyGraph) selectAvailableSteps() []string {
	options := make([]string, 0)
	for nodeID, dependancyIDs := range dg {
		if len(dependancyIDs) == 0 {
			options = append(options, nodeID)
		}
	}
	sort.Strings(options)
	return options
}

func (dg dependancyGraph) removeProcessedStep(nodeIDToDelete string) {
	for nodeID, dependancyIDs := range dg {
		if nodeID == nodeIDToDelete {
			delete(dg, nodeID)
			continue
		}

		// Remove node from dependancy of others
		newDependancies := make([]string, 0)
		for _, d := range dependancyIDs {
			if d != nodeIDToDelete {
				newDependancies = append(newDependancies, d)
			}
		}
		dg[nodeID] = newDependancies
	}
}
