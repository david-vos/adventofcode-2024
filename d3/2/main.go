package main

import (
	"fmt"
	"regexp"
	"strconv"
  "os"
  "bufio"
)


func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	
	if err := scanner.Err(); err != nil {
		return "", err
	}
	
	return content, nil
}

func main() {
  filePath := "input.txt"
	input, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	pattern := `do\(\)|don't\(\)|mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	output := 0
	enabled := true

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			if len(match) > 2 {
				num1Str := match[1]
				num2Str := match[2]

				num1, err1 := strconv.Atoi(num1Str)
				num2, err2 := strconv.Atoi(num2Str)

				if err1 != nil || err2 != nil {
					fmt.Println("Error parsing numbers:", err1, err2)
					continue
				}

				output += num1 * num2
			}
		}
	}

	fmt.Println(output)
}


