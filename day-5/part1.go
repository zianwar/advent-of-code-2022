package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var m = map[int][]string{
	1: []string{"S", "T", "H", "F", "W", "R"},
	2: []string{"S", "G", "D", "Q", "W"},
	3: []string{"B", "T", "W"},
	4: []string{"D", "R", "W", "T", "N", "Q", "Z", "J"},
	5: []string{"F", "B", "H", "G", "L", "V", "T", "Z"},
	6: []string{"L", "P", "T", "C", "V", "B", "S", "G"},
	7: []string{"Z", "B", "R", "T", "W", "G", "P"},
	8: []string{"N", "G", "M", "T", "C", "J", "R"},
	9: []string{"L", "G", "B", "W"},
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

	// Consume the first useless part
	for i := 0; i < 10; i++ {
		scanner.Scan()
	}

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.ReplaceAll(line, "move ", "")
		line = strings.ReplaceAll(line, " from ", ",")
		line = strings.ReplaceAll(line, " to ", ",")

		parts := strings.Split(line, ",")
		count, from, to := toInt(parts[0]), toInt(parts[1]), toInt(parts[2])

		if len(m[from]) < count {
			count = len(m[from])
		}

		for i := 0; i < count; i++ {
			fromList, toList := m[from], m[to]
			fromListLast := fromList[len(fromList)-1]

			fromList = fromList[:len(fromList)-1]
			toList = append(toList, fromListLast)

			m[to] = toList
			m[from] = fromList
		}

	}

	for i := 1; i <= 9; i++ {
		v := m[i]
		if len(v) > 0 {
			fmt.Print(v[len(v)-1])
		}
	}

}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
