package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

var directions = []Position{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return input
}

func findGuardPosition(inputMap [][]rune) (Position, bool) {
	for i, row := range inputMap {
		for j, cell := range row {
			if cell == '^' {
				inputMap[i][j] = '.' // Replace guard's initial position with empty space
				return Position{i, j}, true
			}
		}
	}
	return Position{}, false
}

func simulateGuardMovement(inputMap [][]rune) int {
	guardPos, found := findGuardPosition(inputMap)
	if !found {
		panic("Guard position not found")
	}
	guardDir := 0 // Initially facing up

	visited := make(map[Position]struct{})
	visited[guardPos] = struct{}{}

	for {
		nextPos := Position{
			x: guardPos.x + directions[guardDir].x,
			y: guardPos.y + directions[guardDir].y,
		}

		// Check if the guard has left the mapped area
		if nextPos.x < 0 || nextPos.x >= len(inputMap) || nextPos.y < 0 || nextPos.y >= len(inputMap[0]) {
			break
		}

		if inputMap[nextPos.x][nextPos.y] == '#' {
			// Turn right 90 degrees
			guardDir = (guardDir + 1) % 4
		} else {
			// Move forward
			guardPos = nextPos
			visited[guardPos] = struct{}{}
		}
	}

	return len(visited)
}

func main() {
	inputMap := readInput("d6Input.txt")
	distinctPositionsVisited := simulateGuardMovement(inputMap)
	fmt.Printf("Distinct positions visited: %d\n", distinctPositionsVisited)
}
