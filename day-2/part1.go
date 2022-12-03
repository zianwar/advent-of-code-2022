package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Hand string

func (h Hand) Rock() bool {
	return h == "A" || h == "X"
}
func (h Hand) Paper() bool {
	return h == "B" || h == "Y"
}
func (h Hand) Scissors() bool {
	return h == "C" || h == "Z"
}

func (h Hand) CalculateScore(opponent Hand) int {
	winScore := 0
	if h.Rock() {
		if opponent.Scissors() {
			winScore = 6
		} else if opponent.Paper() {
			winScore = 0
		} else {
			winScore = 3
		}
	} else if h.Paper() {
		if opponent.Rock() {
			winScore = 6
		} else if opponent.Scissors() {
			winScore = 0
		} else {
			winScore = 3
		}
	} else if h.Scissors() {
		if opponent.Paper() {
			winScore = 6
		} else if opponent.Rock() {
			winScore = 0
		} else {
			winScore = 3
		}
	}
	return SelectionScore(h) + winScore
}

func SelectionScore(h Hand) int {
	if h.Rock() {
		return 1
	}
	if h.Paper() {
		return 2
	}
	if h.Scissors() {
		return 3
	}
	return 0
}

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

	totalScore := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, " ")
		opponentHand, myHand := Hand(parts[0]), Hand(parts[1])

		score := myHand.CalculateScore(opponentHand)
		totalScore += score
	}

	fmt.Println(totalScore)
}
