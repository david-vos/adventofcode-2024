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
  return -1
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
  //
  
  output := 0
	for _, row := range inputData {
  

    //get every rule for update
    var matchedRulesForRow [][]int
    for i := 0; i < len(rulesData); i++ {
      rule := rulesData[i]
      if contains(row, rule[0]) && contains(row, rule[1]) {
        matchedRulesForRow = append(matchedRulesForRow, rule)
      }
    }


    hasChange := true
    //main sorting
    for hasChange {
      for i := 0; i < len(matchedRulesForRow); i++ {
        pageRule := matchedRulesForRow[i]
        pageXpos := getIndexOf(row, pageRule[0])
        pageYpos := getIndexOf(row, pageRule[1])
        
        if pageXpos > pageYpos {
          row = []int{0,0}
        }



      // sorts all of them because I missread the question...
      //  if(pageXpos > pageYpos) {
      //    hasChange = false
      //    continue
      //  }

      //  pageXval := row[pageXpos]
      // pageYval := row[pageYpos]

      // row[pageXpos] = pageYval
      // row[pageYpos] = pageXval
      }
      // yeah no need to this it is just there for the sorting...
      hasChange = false
    }



    fmt.Printf("%v\n", row)
    output += row[len(row)/2]
  }
  println(output)
}

