package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
}

func formatInputIntoArray(filePath string) [][]rune {
	file, _ := openFile(filePath)
	defer file.Close()

	var output [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, []rune(line))
	}
	return output
}

func getCurrentDirOfGuard(data [][]rune) [][]int {
	for i, row := range data {
		for j, char := range row {
			switch char {
			case '^':
				return [][]int{
					{-1, 0},
					{i, j}}
			case 'v':
				return [][]int{
					{1, 0},
					{i, j}}
			case '<':
				return [][]int{
					{0, 1},
					{i, j}}
			case '>':
				return [][]int{
					{0, -1},
					{i, j}}
			}
		}
	}
	return [][]int{{0, 0}, {-1, -1}}
}

func getNextGuardPosition(currentPos []int, direction []int) []int {
	return []int{currentPos[0] + direction[0], currentPos[1] + direction[1]}
}

func validateEndOfGame(data [][]rune, nextPos []int) bool {
	if nextPos[0] < 0 || nextPos[0] >= len(data) {
		return true
	}
	if nextPos[1] < 0 || nextPos[1] >= len(data[0]) {
		return true
	}
	return false
}

func checkForOpsticals(data [][]rune, nextPos []int) bool {
	if data[nextPos[0]][nextPos[1]] == '#' {
		return true
	}
	return false
}

func rotateGuard(data [][]rune, currentGuardRotation []int, currentGuardPos []int) [][]rune {
	if currentGuardRotation[0] == 1 {
		data[currentGuardPos[0]][currentGuardPos[1]] = '>'
	} else if currentGuardRotation[0] == -1 {
		data[currentGuardPos[0]][currentGuardPos[1]] = '<'
	} else if currentGuardRotation[1] == 1 {
		data[currentGuardPos[0]][currentGuardPos[1]] = 'v'
	} else if currentGuardRotation[1] == -1 {
		data[currentGuardPos[0]][currentGuardPos[1]] = '^'
	}
	return data
}

func moveTheGuard(data [][]rune, nextPos []int, currentPos []int) [][]rune {
	currentGuard := data[currentPos[0]][currentPos[1]]
	data[currentPos[0]][currentPos[1]] = 'X'
	data[nextPos[0]][nextPos[1]] = currentGuard
	return data
}

func printGrid(data [][]rune) {
	for _, row := range data {
		fmt.Println(string(row))
	}
	fmt.Println("----------------------------------")
}

func getAllGuardPos(data [][]rune) [][]int {
	var guardPos [][]int
	for i, row := range data {
		for j, char := range row {
			if char == 'X' || char == '<' || char == '>' || char == 'v' || char == '^' {
				guardPos = append(guardPos, []int{i, j})
			}
		}
	}
	return guardPos
}

func getPosWeNeedToCheck(data [][]rune) [][]int {
	gaurdIsInPlay := true
	for gaurdIsInPlay {
		dirAndPos := getCurrentDirOfGuard(data)
		direction := dirAndPos[0]
		currentPos := dirAndPos[1]
		nextPos := getNextGuardPosition(currentPos, direction)
		if validateEndOfGame(data, nextPos) {
			gaurdIsInPlay = false
			continue
		}
		if checkForOpsticals(data, nextPos) {
			data = rotateGuard(data, direction, currentPos)
			continue
		}

		moveTheGuard(data, nextPos, currentPos)
	}
	return getAllGuardPos(data)
}

func validateIfTheGuardGotStuck(data [][]rune) bool {
	moveCount := 0
	gaurdIsInPlay := true
	for gaurdIsInPlay {
		if moveCount > len(data)*len(data) {
			return true
		}
		moveCount++
		dirAndPos := getCurrentDirOfGuard(data)
		direction := dirAndPos[0]
		currentPos := dirAndPos[1]
		nextPos := getNextGuardPosition(currentPos, direction)
		if validateEndOfGame(data, nextPos) {
			gaurdIsInPlay = false
			continue
		}
		if checkForOpsticals(data, nextPos) {
			data = rotateGuard(data, direction, currentPos)
			continue
		}

		moveTheGuard(data, nextPos, currentPos)
	}
	return false
}

func main() {
	data := formatInputIntoArray("input.txt")
	amountTheGuardGetsStuck := 0
	possibleObstical := getPosWeNeedToCheck(data)
	fmt.Println(len(possibleObstical))
	for i, pos := range possibleObstical {
		fmt.Println(i)
		data = formatInputIntoArray("input.txt")
		data[pos[0]][pos[1]] = '#'
		if validateIfTheGuardGotStuck(data) {
			amountTheGuardGetsStuck++
		}
	}
	fmt.Println(amountTheGuardGetsStuck)
}
