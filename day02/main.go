// 1) Iterate through each row
// 2) Check for report safety
//   - Must be either increasing or decreasing
//   - Difference between adjacent values at least 1 at most 3
// 3) Solution is number of safe reports

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	records := processInput("input.txt")

	numSafeRecords := 0
	for _, record := range records {

		slopeTrend := 0
		var prevLevel int
		recordSafe := true

		fmt.Println("Record: ", record)
		for i, level := range record {
			if i == 0 {
				prevLevel = level
				continue
			}

			difference := level - prevLevel

			var slope int
			if difference > 0 {
				slope = 1
			} else {
				slope = -1
			}

			if difference < 0 {
				difference *= -1
			}

			if difference > 3 || difference < 1 {
				recordSafe = false
				println("Record unsafe because of difference ", difference, " between ", level, " and ", prevLevel)
				prevLevel = level
				break // Go to next record
			}

			if i == 1 {
				slopeTrend = slope
				prevLevel = level
				continue
			}

			if slope != slopeTrend {
				recordSafe = false
				println("Record unsafe because of slope between ", level, " and ", prevLevel)
				prevLevel = level
				break
			}

			prevLevel = level
		}

		if recordSafe {
			println("Record safe")
			numSafeRecords++
		}
	}

	println("Solution: ", numSafeRecords)
}

func processInput(filePath string) [][]int {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	output := [][]int{}
	for s := range strings.Lines(input) {

		record := []int{}
		for v := range strings.FieldsSeq(s) {
			intVal, _ := strconv.Atoi(v)
			record = append(record, intVal)
		}
		output = append(output, record)
	}

	return output
}
