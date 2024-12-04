package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func extractAndMultiply(filename string) (int, error) {
	var results []int
	mulPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	mulEnabled := true
	dataStr := string(data)

	for i := 0; i < len(dataStr); i++ {
		if len(dataStr[i:]) >= 4 && dataStr[i:i+4] == "do()" {
			mulEnabled = true
			i += 3
		} else if len(dataStr[i:]) >= 7 && dataStr[i:i+7] == "don't()" {
			mulEnabled = false
			i += 6
		}

		if mulEnabled {
			mulPattern.Longest()
			match := mulPattern.FindStringSubmatchIndex(dataStr[i:])
			if match != nil && match[0] == 0 {
				n1, _ := strconv.Atoi(dataStr[i+match[2] : i+match[3]])
				n2, _ := strconv.Atoi(dataStr[i+match[4] : i+match[5]])
				results = append(results, n1*n2)
				i += match[1] - 1
			}
		}
	}

	totalSum := 0
	for _, result := range results {
		totalSum += result
	}

	return totalSum, nil
}

func main() {
	totalSum, err := extractAndMultiply("d3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Total Sum:", totalSum)
}
