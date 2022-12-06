package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(input)

	step := 14
	i := 0
	for i < len(str)-step-1 {
		if !hasDupe(str[i : i+step]) {
			fmt.Println(i + step)
			break
		}
		i++
	}
}

func hasDupe(s string) bool {
	m := make(map[rune]bool)
	for _, c := range s {
		if m[c] {
			return true
		}
		m[c] = true
	}
	return false
}
