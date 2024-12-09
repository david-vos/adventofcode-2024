package main

import (
	"bufio"
	"fmt"
	"os"
)

// GetGrid reads the grid from a file and returns it as a 2D array of runes.
func getGrid() ([][]rune, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return grid, nil
}

func countXMasShapes(grid [][]rune, i, j int) int {
	if i-1 < 0 || i+1 >= len(grid) || j-1 < 0 || j+1 >= len(grid[0]) {
		return 0
	}

	pattern := []rune{
		grid[i-1][j-1], // Top-left
		grid[i-1][j+1], // Top-right
		grid[i+1][j-1], // Bottom-left
		grid[i+1][j+1], // Bottom-right
		grid[i][j],     // Center
	}

	if pattern[4] != 'A' {
		return 0
	}

	if (pattern[0] == 'M' && pattern[1] == 'S' && pattern[2] == 'M' && pattern[3] == 'S') || // M.S / S.M pattern
		(pattern[0] == 'S' && pattern[1] == 'M' && pattern[2] == 'S' && pattern[3] == 'M') || // S.M / M.S pattern
		(pattern[0] == 'M' && pattern[1] == 'M' && pattern[2] == 'S' && pattern[3] == 'S') || // M.M / S.S
		(pattern[0] == 'S' && pattern[1] == 'S' && pattern[2] == 'M' && pattern[3] == 'M') { // S.S / M.M 
		return 1
	}

	return 0
}

func main() {
	grid, err := getGrid()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalCount := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			totalCount += countXMasShapes(grid, i, j)
		}
	}

	fmt.Printf("Total unique 'X-MAS' shapes found: %d\n", totalCount)
}

