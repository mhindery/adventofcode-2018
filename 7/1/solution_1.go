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
	// fmt.Println("---------------")

	sequence := ""

	// fmt.Println(nodeDependancies)

	for len(nodeDependancies) > 0 {
		n := nodeDependancies.selectAvailableStep()
		sequence += n
		// fmt.Println(n)
		nodeDependancies.removeProcessedStep(n)
		// fmt.Println(nodeDependancies)
	}
	fmt.Println(sequence)

}

// selectAvailableStep selects the next step todo: alphabetically first from those without dependancies
func (dg dependancyGraph) selectAvailableStep() string {
	options := make([]string, 0)
	for nodeID, dependancyIDs := range dg {
		if len(dependancyIDs) == 0 {
			options = append(options, nodeID)
		}
	}
	sort.Strings(options)
	return options[0]
}

func (dg dependancyGraph) removeProcessedStep(nodeIDToDelete string) {
	for nodeID, dependancyIDs := range dg {
		if nodeID == nodeIDToDelete {
			delete(dg, nodeID)
			continue
		}

		newDependancies := make([]string, 0)
		for _, d := range dependancyIDs {
			if d != nodeIDToDelete {
				newDependancies = append(newDependancies, d)
			}
		}
		dg[nodeID] = newDependancies
	}
}
