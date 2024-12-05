package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
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
				fmt.Println("Error converting to integer:", err)
				return
			}
			row = append(row, number)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}


safeCount := 0

for _, row := range grid {
    if len(row) < 2 {
        continue
    }
    direction := row[0] < row[1]
    isSafe := true
    for j := 0; j < len(row)-1; j++ {
        diff := row[j+1] - row[j]
        if diff < -3 || diff > 3 || diff == 0 {
            isSafe = false
            break
        }
        if (direction && diff < 0) || (!direction && diff > 0) {
            isSafe = false
            break
        }
    }
    if isSafe {
        safeCount++
    }
}

  println(safeCount)
}


