package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**************
Create per day an array of 120 minutes to indicate state
Store arrays in map with key guardID-day


**************/

func actionToGuardID(s string) string {
	parts := strings.Split(s, " ")
	return parts[1][1:]
}

func timestampToIndex(s string) int {
	parts := strings.Split(s, ":")
	minutes, _ := strconv.Atoi(parts[1])
	return minutes
}

func extractInfoLine(s string) (string, string, int, string) {
	// timing := s[1:17]
	timingParts := strings.Split(s[1:17], " ")
	date := timingParts[0]
	timestamp := timingParts[1]
	action := s[19:]
	index := timestampToIndex(timestamp)

	return date, timestamp, index, action
}

func markSleeping(schedule map[string][]int, id string, from int, to int) {
	for index := from; index < to; index++ {
		schedule[id][index]++
	}
}

// combines all arrays of a guard
func calculatePerGuard(schedule map[string][]int) map[string][]int {
	sums := make(map[string][]int)
	for id, arr := range schedule {
		guardID := strings.Split(id, "-")[3]
		if _, ok := sums[guardID]; !ok {
			sums[guardID] = make([]int, 60, 60)
		}
		for index, value := range arr {
			sums[guardID][index] += value
		}
	}
	return sums
}

func calculateStats(schedule map[string][]int) map[string][]int {
	// guardID: [totalSleepMinutes worstMinuteCount worstMinute]
	statsMap := make(map[string][]int)

	for k, arr := range schedule {
		statsMap[k] = make([]int, 3, 3)

		sum := 0

		maxValue := 0
		maxValueMinute := 0

		for v := range arr {
			sum += arr[v]

			if arr[v] > maxValue {
				maxValue = arr[v]
				maxValueMinute = v
			}
		}

		statsMap[k][0] = sum
		statsMap[k][1] = maxValue
		statsMap[k][2] = maxValueMinute
	}

	// fmt.Println("GuardID\t\tSum\tCount\tMinute")
	// for k, v := range statsMap {
	// 	fmt.Printf("%v\t\t%v\t%v\t%v\n", k, v[0], v[1], v[2])
	// }

	return statsMap
}

func idToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

var (
	previousIndex = 0
)

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

	sort.Strings(inputs)

	scheduler := make(map[string][]int)
	arrayID := ""
	// previousAction := ""
	previousIndex := 999999
	// _ = previousAction

	// Fill in schedule per day
	for _, line := range inputs {
		date, timestamp, index, action := extractInfoLine(line)
		_ = timestamp

		// New day: reset
		if strings.Contains(action, "begins shift") {
			arrayID = date + "-" + actionToGuardID(action)
			scheduler[arrayID] = make([]int, 60, 60)
			// previousAction = action
			previousIndex = index
			continue
		}

		if action == "wakes up" {
			markSleeping(scheduler, arrayID, previousIndex, index)
		}

		// previousAction = action
		previousIndex = index
	}

	// Total sleeps
	totalSchedule := calculatePerGuard(scheduler)

	// Print per-guard status
	statsMap := calculateStats(totalSchedule)

	// find entry with largest value in first column
	id := ""
	value := 0

	for k, v := range statsMap {
		if v[0] > value {
			value = v[0]
			id = k
		}
	}
	fmt.Println(id, statsMap[id], "=>", idToInt(id)*statsMap[id][2])
}
