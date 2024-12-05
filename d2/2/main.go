package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(row []int) bool {
	if len(row) < 2 {
		return true
	}

	direction := row[0] < row[1]
	for i := 0; i < len(row)-1; i++ {
		diff := row[i+1] - row[i]
		if diff < -3 || diff > 3 || diff == 0 || (direction && diff < 0) || (!direction && diff > 0) {
			return false
		}
	}
	return true
}

func readFileToGrid(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var row []int
		for _, field := range fields {
			number, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("error converting to integer: %w", err)
			}
			row = append(row, number)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return grid, nil
}

func main() {
	grid, err := readFileToGrid("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	safeCount := 0

	for _, row := range grid {
		if isSafe(row) {
			safeCount++
			continue
		}

		for i := 0; i < len(row); i++ {
			modifiedRow := []int{}
			for j := 0; j < len(row); j++ {
				if j != i {
					modifiedRow = append(modifiedRow, row[j])
				}
			}

			if isSafe(modifiedRow) {
				safeCount++
				break 
			}
		}
	}

	fmt.Println(safeCount)
}



