package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
  "sort"
)

func readFileToArrays(filename string) ([]int, []int, error) {
	const size = 1000
	array1 := make([]int, 0, size)
	array2 := make([]int, 0, size)

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(array1) >= size {
			break
		}
		parts := strings.Fields(scanner.Text())
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	return array1, array2, scanner.Err()
}

func main() {
	filename := "input.txt"
	array1, array2, err := readFileToArrays(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

  sort.Ints(array1)
  sort.Ints(array2)

  totalDif := 0;
  for i := 0; i < len(array1); i++ { 
    if(array1[i] >= array2[i]) {
      totalDif += array1[i] - array2[i]
    } else {
      totalDif += array2[i] - array1[i]
    }
  }

 fmt.Println(totalDif)
}

