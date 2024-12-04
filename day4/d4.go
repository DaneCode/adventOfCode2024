package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readWordSearch(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var grid []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        grid = append(grid, strings.TrimSpace(scanner.Text()))
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return grid, nil
}

func countWordOccurrences(grid []string, word string) int {
    rows, cols := len(grid), len(grid[0])
    wordLen := len(word)
    count := 0

    checkDirection := func(x, y, dx, dy int) bool {
        for i := 0; i < wordLen; i++ {
            nx, ny := x+i*dx, y+i*dy
            if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != word[i] {
                return false
            }
        }
        return true
    }

    directions := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}}

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == word[0] {
                for _, dir := range directions {
                    if checkDirection(i, j, dir[0], dir[1]) {
                        count++
                    }
                }
            }
        }
    }

    return count
}

func main() {
    grid, err := readWordSearch("day4Input.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    word := "XMAS"
    occurrences := countWordOccurrences(grid, word)
    fmt.Printf("Number of occurrences of the word \"%s\": %d\n", word, occurrences)
}