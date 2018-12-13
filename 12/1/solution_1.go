package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func getTransition(s string) (string, string) {
	parts := strings.Split(s, " => ")
	return parts[0], parts[1]
}

type pot struct {
	state string
	id    int
}

func printPots(l *list.List) {
	s := ""
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(pot).id < -3 {
			continue
		}
		s += e.Value.(pot).state
	}
	fmt.Println(s)
}

func sumPots(l *list.List) (sum int) {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(pot).state == "#" {
			sum += e.Value.(pot).id
		}
	}
	return sum
}

func getState(e *list.Element) string {
	s := e.Value.(pot).state
	s = s + e.Next().Value.(pot).state
	s = s + e.Next().Next().Value.(pot).state
	s = e.Prev().Value.(pot).state + s
	s = e.Prev().Prev().Value.(pot).state + s
	return s
}

func doGeneration(l *list.List, transitions map[string]string) *list.List {
	new := list.New()
	for e := l.Front().Next().Next(); e != l.Back().Prev(); e = e.Next() {
		p := e.Value.(pot)
		state := getState(e)
		newState, ok := transitions[state]
		if !ok {
			newState = "."
		}
		// fmt.Println("\t", p.id, ":", state, "=>", newState)
		new.PushBack(pot{
			state: newState,
			id:    p.id,
		})
	}
	return new
}

func padEmpty(l *list.List) {
	// add empty pots at the beginning and end
	count := 3
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(pot)
		if p.state == "#" {
			break
		}
		count--
	}
	for count > 0 {
		p := l.Front().Value.(pot)
		l.PushFront(pot{
			state: ".",
			id:    p.id - 1,
		})
		count--
	}

	count = 3
	for e := l.Back(); e != nil; e = e.Prev() {
		p := e.Value.(pot)
		if p.state == "#" {
			break
		}
		count--
	}
	for count > 0 {
		p := l.Back().Value.(pot)
		l.PushBack(pot{
			state: ".",
			id:    p.id + 1,
		})
		count--
	}

}

func initalizeList(s string) *list.List {
	s = strings.TrimPrefix(s, "initial state: ")
	l := list.New()

	id := 0
	for i := 0; i < len(s); i++ {
		pot := pot{
			state: string(s[i]),
			id:    id,
		}
		l.PushBack(pot)
		id++
	}

	return l
}

func main() {

	inputs := readInput("../real_input.txt")

	pots := initalizeList(inputs[0])
	padEmpty(pots)

	transitions := make(map[string]string)
	for _, line := range inputs[1:] {
		if len(line) > 0 {
			in, out := getTransition(line)
			transitions[in] = out
		}
	}

	fmt.Print(0, ": \t")
	printPots(pots)
	for it := 1; it < 21; it++ {
		pots = doGeneration(pots, transitions)
		padEmpty(pots)
		fmt.Print(it, ": \t")
		printPots(pots)
	}

	fmt.Println(sumPots(pots))

}
