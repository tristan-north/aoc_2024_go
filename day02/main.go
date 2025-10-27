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

	// Part One
	println("Part One")
	numSafeRecordsOne := 0
	for _, record := range records {
		if isRecordSafe(record) {
			numSafeRecordsOne++
		}
	}

	// Part Two
	println("Part Two")
	numSafeRecordsTwo := 0
	for _, record := range records {
		recordSafe := false
		for i := range record {
			mutatedRecord := []int{}
			mutatedRecord = append(mutatedRecord, record[:i]...)
			mutatedRecord = append(mutatedRecord, record[i+1:]...)
			if isRecordSafe(mutatedRecord) {
				recordSafe = true
				break
			}
		}

		if recordSafe {
			numSafeRecordsTwo++
		}
	}

	println("Solution Part One: ", numSafeRecordsOne)
	println("Solution Part Two: ", numSafeRecordsTwo)
}

func isRecordSafe(record []int) bool {
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
		differenceAbs := difference
		if differenceAbs < 0 {
			differenceAbs *= -1
		}

		prevLevel = level

		if differenceAbs > 3 || differenceAbs < 1 {
			recordSafe = false
			println("Record unsafe because of difference ", differenceAbs, " between ", level, " and ", prevLevel)
			break // Go to next record
		}

		var slope int
		if difference > 0 {
			slope = 1
		} else {
			slope = -1
		}

		if i == 1 {
			slopeTrend = slope
			continue
		}

		if slope != slopeTrend {
			recordSafe = false
			println("Record unsafe because of slope between ", level, " and ", prevLevel)
			break
		}

	}

	return recordSafe
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
