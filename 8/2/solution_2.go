package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	nodeID      string
	numChildren int
	numMetadata int
	children    []*node
	metadata    []int

	parent *node
}

func infoFromLine(s string) []int {
	splitted := strings.Split(s, " ")
	asInts := make([]int, 0, len(splitted))
	for _, stringValue := range splitted {
		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			log.Fatal(err)
		}
		asInts = append(asInts, intValue)
	}
	return asInts
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

var (
	nodeID byte = 65 // use for easy representation
)

func constructNodeTree(sequence []int, parentNode *node) ([]int, *node) {
	nodeInfo := sequence[:2]
	sequence = sequence[2:]

	n := &node{
		nodeID:      string(nodeID),
		numChildren: nodeInfo[0],
		numMetadata: nodeInfo[1],
		children:    make([]*node, 0, nodeInfo[0]),
		metadata:    make([]int, 0, nodeInfo[1]),
		parent:      parentNode,
	}
	nodeID++

	for index := 0; index < n.numChildren; index++ {
		trimmedSequence, child := constructNodeTree(sequence, n)
		sequence = trimmedSequence
		n.children = append(n.children, child)
	}

	n.metadata = sequence[:n.numMetadata]
	sequence = sequence[n.numMetadata:]

	// fmt.Println(nodeInfo, sequence)
	return sequence, n
}

func printTree(n *node, sep string) {
	fmt.Printf("%v%v\n", sep, n)
	for _, child := range n.children {
		printTree(child, sep+"  ")
	}
}

func getValue(n *node) int {
	// fmt.Println("calculating", n)
	value := 0
	if n.numChildren == 0 {
		for _, val := range n.metadata {
			value += val
		}
	} else {
		for _, index := range n.metadata {
			if index > n.numChildren {
				continue
			}
			value += getValue(n.children[index-1]) // indexing 0-based <> file-format
		}
	}
	return value
}

func main() {

	inputs := readInput("../real_input.txt")
	sequence := infoFromLine(inputs[0])
	// fmt.Println(sequence)
	_, rootNode := constructNodeTree(sequence, nil)
	printTree(rootNode, "")
	fmt.Println(getValue(rootNode))
}
