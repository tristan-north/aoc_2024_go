// Part One
// 1) Sort each list
// 2) Get the difference between corresponding values
// 3) Add up all the differences

// Part Two
// 1) For each number in left column, find number of occurances in right column
// 2) Increase result by left column value * num times in right col

package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	left, right := processInput("input.txt")

	// Part One
	accum := 0
	for i := range left {
		accum += abs(left[i] - right[i])
	}
	println("Solution Part One: ", accum)

	// Part Two
	accum = 0
	var numOccurances int
	for _, v := range left {

		numOccurances = 0
		for _, k := range right {
			if v == k {
				numOccurances++
			}
		}

		accum += v * numOccurances
	}

	println("Solution Part Two: ", accum)
}

func processInput(filePath string) ([]int, []int) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	left := []int{}
	right := []int{}

	for line := range strings.Lines(input) {
		fields := strings.Fields(line)

		leftInt, _ := strconv.Atoi(fields[0])
		rightInt, _ := strconv.Atoi(fields[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
