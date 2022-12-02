package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// PriorityQueue implementation
type PQ []int

func (pq *PQ) Len() int {
	return len(*pq)
}
func (pq PQ) Less(i int, j int) bool {
	return pq[i] > pq[j]
}
func (pq PQ) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(v interface{}) {
	*pq = append(*pq, v.(int))
}
func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	last := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return last
}

func main() {
	pq := &PQ{}
	heap.Init(pq)

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elfCalories := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			heap.Push(pq, elfCalories)
			elfCalories = 0
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		elfCalories += n
	}

	// reached the end, count the last Elf's calories
	heap.Push(pq, elfCalories)

	// Count total calories for the top 3 Elves
	totalCalories := 0
	for i := 0; i < 3; i++ {
		n := heap.Pop(pq).(int)
		totalCalories += n
	}

	fmt.Println(totalCalories)
}
