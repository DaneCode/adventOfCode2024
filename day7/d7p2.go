package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func evaluateExpression(numbers []int, operators []string) int {
	result := numbers[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result += numbers[i+1]
		case "*":
			result *= numbers[i+1]
		case "||":
			result, _ = strconv.Atoi(fmt.Sprintf("%d%d", result, numbers[i+1]))
		}
	}
	return result
}

func canBeTrue(testValue int, numbers []int) bool {
	if len(numbers) == 1 {
		return numbers[0] == testValue
	}

	operators := []string{"+", "*", "||"}
	combinations := getCombinations(operators, len(numbers)-1)

	for _, ops := range combinations {
		if evaluateExpression(numbers, ops) == testValue {
			return true
		}
	}
	return false
}

func getCombinations(operators []string, length int) [][]string {
	if length == 1 {
		combinations := make([][]string, len(operators))
		for i, op := range operators {
			combinations[i] = []string{op}
		}
		return combinations
	}

	smallerCombinations := getCombinations(operators, length-1)
	combinations := [][]string{}

	for _, op := range operators {
		for _, smallerCombination := range smallerCombinations {
			combo := append([]string{op}, smallerCombination...)
			combinations = append(combinations, combo)
		}
	}

	return combinations
}

func main() {
	totalCalibrationResult := 0

	data, err := ioutil.ReadFile("d7Input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, ":")
		testValue, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numberStrings := strings.Fields(strings.TrimSpace(parts[1]))
		numbers := make([]int, len(numberStrings))
		for i, numStr := range numberStrings {
			numbers[i], _ = strconv.Atoi(numStr)
		}

		if canBeTrue(testValue, numbers) {
			totalCalibrationResult += testValue
		}
	}

	fmt.Println(totalCalibrationResult)
}
