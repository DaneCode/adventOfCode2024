package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Step 1: Read the data from d1Input.txt
	file, err := os.Open("d1Input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var array1, array2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line) // Use Fields to handle multiple spaces
		if len(nums) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}
		num1, err1 := strconv.Atoi(nums[0])
		num2, err2 := strconv.Atoi(nums[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting numbers:", nums[0], nums[1])
			continue
		}
		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Step 3: Sort both slices
	sort.Ints(array1)
	sort.Ints(array2)

	// Step 4: Calculate the absolute difference between corresponding elements of the two slices
	var array3 []int
	for i := range array1 {
		diff := array1[i] - array2[i]
		if diff < 0 {
			diff = -diff
		}
		array3 = append(array3, diff)
	}

	// Step 5: Sum all the values in the third slice
	totalDifference := 0
	for _, value := range array3 {
		totalDifference += value
	}

	fmt.Println("Total difference:", totalDifference)
}
