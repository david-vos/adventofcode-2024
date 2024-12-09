package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
)


func getGrid() ( [][]rune, error)  {
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
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
  return grid, nil

}

// dir vects same I justed for ray tracing :eyes:
var directions = [][]int{
	{0, 1},  // Right (horizontal)
	{1, 0},  // Down (vertical)
	{1, 1},  // Diagonal down-right
	{1, -1}, // Diagonal down-left
}

var visited = make(map[string]bool)

// Again thx GPT for telling me how to do this...
func createKey(positions [][2]int) string {
	var keyParts []string
	for _, pos := range positions {
		keyParts = append(keyParts, fmt.Sprintf("%d-%d", pos[0], pos[1]))
	}
	return strings.Join(keyParts, ",")
}

func countXmas(grid [][]rune, i, j int) int {
	count := 0

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		chars := []rune{}
		var positions [][2]int

		for step := 0; step < 4; step++ {
			ni, nj := i+(dx*step), j+(dy*step)
			if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
				break
			}
			chars = append(chars, grid[ni][nj])
			positions = append(positions, [2]int{ni, nj})
		}

		if len(chars) == 4 {
			str := string(chars)
			if str == "XMAS" || str == "SAMX" {
				key := createKey(positions)
				if !visited[key] {
					visited[key] = true
					count++
				}
			}
		}
	}

	return count
}

func main() {
	
  grid, err := getGrid()
  if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			totalCount += countXmas(grid, i, j)
		}
	}

	fmt.Printf("Total unique 'XMAS' or 'SAMX' found: %d\n", totalCount) 
}

