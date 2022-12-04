package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ",")
		p1, p2 := parts[0], parts[1]

		p1Parts, p2Parts := strings.Split(p1, "-"), strings.Split(p2, "-")
		p1s, p1e := toInt(p1Parts[0]), toInt(p1Parts[1])
		p2s, p2e := toInt(p2Parts[0]), toInt(p2Parts[1])

		if (p1s <= p2s && p2e <= p1e) || (p2s <= p1s && p1e <= p2e) ||
			(p1s <= p2s && p2s <= p1e) || (p2s <= p1s && p1s <= p2e) ||
			(p1s <= p2e && p2e <= p1e) || (p2s <= p1e && p1e <= p2e) {
			count += 1
		}

	}

	fmt.Println(count)
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
