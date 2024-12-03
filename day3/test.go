package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func mul(match1 string, match2 string) (int, error) {
	num1, err := strconv.Atoi(match1)
	if err != nil {
		fmt.Println("could not parse int", err)
		return 0, err
	}
	num2, err := strconv.Atoi(match2)
	if err != nil {
		fmt.Println("could not parse int", err)
		return 0, err
	}
	return num1 * num2, nil
}

func part1(fileContents string) (int, error) {
	r, err := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	if err != nil {
		fmt.Println("could not compile regex", err)
	}
	matches := r.FindAllStringSubmatch(fileContents, -1)
	output := 0
	for _, match := range matches {
		num, err := mul(match[1], match[2])
		if err != nil {
			fmt.Println("could not process tokens", err)
			return 0, err
		}
		output += num
	}
	return output, nil
}

func part2(fileContents string) (int, error) {
	r, err := regexp.Compile(`don't\(\)|mul\(([0-9]+),([0-9]+)\)|do\(\)`)
	if err != nil {
		fmt.Println("could not compile regex", err)
	}
	matches := r.FindAllStringSubmatch(fileContents, -1)
	output := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "don't()" {
			enabled = false
		} else if match[0] == "do()" {
			enabled = true
		} else if enabled {
			num, err := mul(match[1], match[2])
			if err != nil {
				fmt.Println("could not process tokens", err)
				return 0, err
			}
			output += num
		}
	}
	return output, nil
}

func main() {
	// fileContent := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	data, err := os.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fileContent := string(data)
	output, err := part1(fileContent)
	if err != nil {
		fmt.Println("Error during part1:", err)
		return
	}
	fmt.Println("results part 1", output)
	output2, err := part2(fileContent)
	if err != nil {
		fmt.Println("Error during part2:", err)
		return
	}
	fmt.Println("results part 2", output2)
}
