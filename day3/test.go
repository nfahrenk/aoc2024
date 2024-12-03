package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string
	fileContent := string(data)

	r, err := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	if err != nil {
		fmt.Println("could not compile regex", err)
	}
	matches := r.FindAllStringSubmatch(fileContent, -1)
	output := 0
	for _, match := range matches {
		fmt.Println(match)
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("could not parse int", err)
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("could not parse int", err)
		}
		output += num1 * num2
	}
	fmt.Println("results", output)
}
