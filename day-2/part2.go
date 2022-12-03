package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type RPS int

const (
	Rock RPS = iota
	Paper
	Scissors
)

type Outcome string

func (o Outcome) Lose() bool {
	return o == "X"
}
func (o Outcome) Draw() bool {
	return o == "Y"
}
func (o Outcome) Win() bool {
	return o == "Z"
}

type Hand string

func (h Hand) Rock() bool {
	return h == "A"
}
func (h Hand) Paper() bool {
	return h == "B"
}
func (h Hand) Scissors() bool {
	return h == "C"
}

func (opponentHand Hand) CalculateScore(outcome Outcome) int {
	selectionScore := 0
	winScore := 0
	if opponentHand.Rock() {
		if outcome.Lose() {
			selectionScore = SelectionScore(Scissors)
			winScore = 0
		} else if outcome.Draw() {
			selectionScore = SelectionScore(Rock)
			winScore = 3
		} else if outcome.Win() {
			selectionScore = SelectionScore(Paper)
			winScore = 6
		}
	} else if opponentHand.Paper() {
		if outcome.Lose() {
			selectionScore = SelectionScore(Rock)
			winScore = 0
		} else if outcome.Draw() {
			selectionScore = SelectionScore(Paper)
			winScore = 3
		} else if outcome.Win() {
			selectionScore = SelectionScore(Scissors)
			winScore = 6
		}
	} else if opponentHand.Scissors() {
		if outcome.Lose() {
			selectionScore = SelectionScore(Paper)
			winScore = 0
		} else if outcome.Draw() {
			selectionScore = SelectionScore(Scissors)
			winScore = 3
		} else if outcome.Win() {
			selectionScore = SelectionScore(Rock)
			winScore = 6
		}
	}
	return selectionScore + winScore
}

func SelectionScore(x RPS) int {
	switch x {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
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
		opponentHand, outcome := Hand(parts[0]), Outcome(parts[1])

		score := opponentHand.CalculateScore(outcome)
		totalScore += score
	}

	fmt.Println(totalScore)
}
