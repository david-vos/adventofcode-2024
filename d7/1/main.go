package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
}

func Split(r rune) bool {
	return r == ':' || r == ' '
}

func readRulesTo2DArray(filePath string) ([][]int, error) {
	file, err := openFile(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, Split)
		var row []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error converting to integer: %v", err)
			}
			row = append(row, num)
		}
		data = append(data, row)
	}
	return data, nil
}

func getAllRandomOperations(row []int) bool {
	// we do -2 on the length because first we should remove the result at index 0
	// and another number because the needed math requires 2 points IE -1
	arr := make([]int, len(row)-2)
	hasValidOption := false
	GenerateCombinations(arr, 0, func(combination []int) {
		if isValid(row, combination) {
			hasValidOption = true
		}
	})
	return hasValidOption

}

func GenerateCombinations(arr []int, index int, callback func([]int)) {
	// Thx chatGPT for helping me with this recursion
	if index == len(arr) {
		callback(arr)
		return
	}

	arr[index] = 1
	GenerateCombinations(arr, index+1, callback)
	arr[index] = 2
	GenerateCombinations(arr, index+1, callback)
}

func isValid(row []int, opList []int) bool {
	output := row[1]
	for i := 1; i < len(row)-1; i++ {
		a := output
		b := row[i+1]

		if opList[i-1] == 1 {
			output = a + b
		} else {
			output = a * b
		}
	}
	if output == row[0] {
		return true
	}
	return false
}

func main() {
	data, err := readRulesTo2DArray("input.txt")
	fmt.Println(err)

	count := 0
	for _, row := range data {
		isValid := getAllRandomOperations(row)
		if isValid {
			count += row[0]
			continue
		}
	}

	fmt.Println(count)
}
