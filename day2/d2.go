package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(report []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff < 1 || diff > 3 {
			increasing = false
		}
		if diff > -1 || diff < -3 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func main() {
	// Step 1: Read the data from d2input.txt
	file, err := os.Open("d2input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		var report []int
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error converting number:", num)
				return
			}
			report = append(report, n)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Step 3-5: Check each report and count the number of safe reports
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports:", safeCount)
}
