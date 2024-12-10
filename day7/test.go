package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	result int
	inputs []int
}

func readInput(filename string) ([]Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []Input{}
	for scanner.Scan() {
		line := scanner.Text()

		colon := strings.Split(line, ":")
		result, err := strconv.Atoi(colon[0])
		if err != nil {
			fmt.Println("could not parse int", err)
			return nil, err
		}
		parts := strings.Split(strings.TrimLeft(colon[1], " "), " ")

		nums := []int{}
		for _, val := range parts {
			parsed, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("could not parse int", err)
				return nil, err
			}
			nums = append(nums, parsed)
		}

		input := Input{inputs: nums, result: result}
		lines = append(lines, input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}

func generateOperators(current string, maxLen int, results *[]string) {
	if len(current) == maxLen {
		*results = append(*results, current)
		return
	}
	for _, op := range []string{"*", "+", "|"} {
		generateOperators(current+op, maxLen, results)
	}
}

func part2(filename string) int {
	inputs, err := readInput(filename)
	if err != nil {
		return 0
	}
	output := 0
	for _, input := range inputs {
		results := make([]string, 0)
		generateOperators("", len(input.inputs)-1, &results)
		for _, combo := range results {
			balance := input.inputs[0]
			for ndx, val := range input.inputs[1:] {
				switch combo[ndx] {
				case '*':
					balance = balance * val
				case '+':
					balance += val
				case '|':
					left := strconv.Itoa(balance)
					right := strconv.Itoa(val)
					balance, err = strconv.Atoi(left + right)
					if err != nil {
						fmt.Println("error during parsing int", err)
						return -1
					}
				default:
					fmt.Println("invalid operator")
					return -1
				}
			}
			if balance == input.result {
				output += input.result
				break
			}
		}
	}
	return output
}

func main() {
	example := part2("./input1.txt")
	if example != 11387 {
		fmt.Println("assertion failed :(", example)
		return
	}
	fmt.Println("assertion passed!!")
	output := part2("./input2.txt")
	fmt.Println("output", output)
}
