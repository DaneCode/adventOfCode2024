package main

import (
	"bufio"
	"fmt"
	"os"
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

	// Step 3: Count the occurrences of each number from the first array in the second array
	counter := make(map[int]int)
	for _, num := range array2 {
		counter[num]++
	}

	// Step 4: Multiply each number from the first array by its count in the second array and store the result in a third array
	var array3 []int
	for _, num := range array1 {
		array3 = append(array3, num*counter[num])
	}

	// Step 5: Sum all the values in the third array to get the total
	total := 0
	for _, value := range array3 {
		total += value
	}

	fmt.Println("Total:", total)
}
