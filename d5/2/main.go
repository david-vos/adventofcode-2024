package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// I dont care that I make this code way better by have just one call for the file and the reader :angry:
func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
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
		parts := strings.Split(line, "|")

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

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return data, nil
}

func readInputTo2DArray(filePath string) ([][]int, error) {
	file, err := openFile(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

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

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return data, nil
}

func contains(s []int, e int) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}


func getIndexOf(s []int, e int) int {
  for index, a := range s {
    if a == e {
      return index
    }
  }
  fmt.Printf("%v\n", s)
  println("seachVal: ", e)
  return -99
}

func main() {
	rulesFilePath := "rules.txt"
  var rulesData [][]int
	rulesData, err := readRulesTo2DArray(rulesFilePath)
	if err != nil {
		fmt.Println("Error reading rules file:", err)
		return
	}

	inputFilePath := "input.txt"
	inputData, err := readInputTo2DArray(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

  // I should prop have used a hash map of some sort
  // for every update
  // I need sleep so this will have to do. Its ugly but I guess it works :shrug:
  
  output := 0
  for _, row := range inputData {

    // Get every rule for update
    var matchedRulesForRow [][]int
    for i := 0; i < len(rulesData); i++ {
        rule := rulesData[i]
        if contains(row, rule[0]) && contains(row, rule[1]) {
            matchedRulesForRow = append(matchedRulesForRow, rule)
        }
    }

    // Check if all rules are satisfied for the row
    alreadyCorrect := true
    for _, rule := range matchedRulesForRow {
        pageXpos := getIndexOf(row, rule[0])
        pageYpos := getIndexOf(row, rule[1])

        if pageXpos > pageYpos {
            alreadyCorrect = false
            break
        }
    }

    if alreadyCorrect {
        // I have no idea of how to delete from an array so I'll just set it to [0,0] 
        row = []int{0, 0}
    } else {
        // Main sorting
        // It took way to long to get the flags for hasBeenSorted correctly.
        // I think this is the best way to do this but there might be a faster way of sorting >,<
        hasBeenSorted := true
        for hasBeenSorted {
            hasBeenSorted = false
            for _, rule := range matchedRulesForRow {
                pageXpos := getIndexOf(row, rule[0])
                pageYpos := getIndexOf(row, rule[1])

                if pageXpos > pageYpos {
                    // fancy def that I saw online :eyes:
                    row[pageXpos], row[pageYpos] = row[pageYpos], row[pageXpos]
                    hasBeenSorted = true
                }
            }
        }
    }

    fmt.Printf("%v\n", row)
    output += row[len(row)/2]
}

  println(output)
}

