// This does not work at the moment and will have to work it out
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func extractAndMultiply(filename string) []int {
	var results []int
	// Updated pattern to strictly match valid mul(X,Y) instructions
	mulPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return results
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mulEnabled := true // Initially, mul instructions are enabled

	for scanner.Scan() {
		line := scanner.Text()
		i := 0
		for i < len(line) {
			// Check for do() and don't() instructions
			if doPattern.MatchString(line[i:]) {
				mulEnabled = true
				i += 4 // Move index past the do() instruction
			} else if dontPattern.MatchString(line[i:]) {
				mulEnabled = false
				i += 7 // Move index past the don't() instruction
			} else if mulEnabled {
				// Check for mul(X,Y) instructions
				matches := mulPattern.FindStringSubmatchIndex(line[i:])
				if matches != nil && matches[0] == 0 {
					n1, _ := strconv.Atoi(line[matches[2]:matches[3]])
					n2, _ := strconv.Atoi(line[matches[4]:matches[5]])
					results = append(results, n1*n2)
					i += matches[1] // Move index past the mul(X,Y) instruction
				} else {
					i++
				}
			} else {
				i++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return results
}

func main() {
	// File path
	filePath := "d3.txt"

	// Get the results
	multiplicationResults := extractAndMultiply(filePath)

	// Calculate the sum of the results
	totalSum := 0
	for _, result := range multiplicationResults {
		totalSum += result
	}

	// Print the sum of the results
	fmt.Println("Sum of multiplication results:")
	fmt.Println(totalSum)
}
