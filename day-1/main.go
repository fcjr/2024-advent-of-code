package main

import (
	"bufio"
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var left []int
	var right []int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	acc := 0
	for i := 0; i < len(left); i++ {
		leftNum := left[i]
		rightNum := right[i]
		var distance int
		if leftNum > rightNum {
			distance = leftNum - rightNum
		} else {
			distance = rightNum - leftNum
		}
		acc += distance
	}
	log.Println(acc)
}
