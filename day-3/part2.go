package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	prioritySum := 0

	i := 0
	group := []string{}
	for scanner.Scan() {
		i++
		line := strings.TrimSpace(scanner.Text())
		group = append(group, line)

		// Process in batch of 3 lines
		if i == 3 {
			prioritySum += groupPriority(group)

			// Reset counter
			i = 0
			group = []string{}
		}
	}

	fmt.Println(prioritySum)
}

func groupPriority(lines []string) int {
	freq := make(map[rune]int)

	for _, line := range lines {
		// Count chars of each line once
		seen := make(map[rune]bool)
		for _, c := range line {
			if _, ok := seen[c]; ok {
				continue
			}
			freq[c]++
			seen[c] = true
		}
	}

	// Count the chars that appeared in all lines
	prio := 0
	for c, v := range freq {
		if v >= len(lines) {
			prio += priority(c)
		}
	}

	return prio
}

func priority(c rune) int {
	if int(c) >= int('a') && int(c) <= int('z') {
		return int(c) - (int('a') - 1)
	}
	if int(c) >= int('A') && int(c) <= int('Z') {
		return int(c) - (int('A') - 1) + 26
	}
	return 0
}
