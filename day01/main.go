// 1) Sort each list
// 2) Get the difference between corresponding values
// 3) Add up all the differences

package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	var left []int
	var right []int

	for line := range strings.Lines(input) {
		fields := strings.Fields(line)

		leftInt, _ := strconv.Atoi(fields[0])
		rightInt, _ := strconv.Atoi(fields[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	slices.Sort(left)
	slices.Sort(right)

	accum := 0
	for i := range left {
		accum += abs(left[i] - right[i])
	}

	println("Solution: ", accum)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
