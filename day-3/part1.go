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

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		left, right := line[:len(line)/2], line[len(line)/2:]

		m := make(map[rune]bool)
		for _, c := range left {
			m[c] = true
		}

		for _, c := range right {
			if _, ok := m[c]; ok {
				prioritySum += priority(c)
				delete(m, c)
			}
		}
	}

	fmt.Println(prioritySum)
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
