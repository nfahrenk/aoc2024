package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Left  int
	Right int
}

func readInput(filename string) ([]Rule, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := []Rule{}
	tests := [][]int{}

	isRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			isRules = false
			continue
		}
		var delim string
		if isRules {
			delim = "|"
		} else {
			delim = ","
		}
		parts := strings.Split(line, delim)

		nums := []int{}
		for _, val := range parts {
			parsed, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("could not parse int", err)
				return nil, nil, err
			}
			nums = append(nums, parsed)
		}

		if isRules {
			rules = append(rules, Rule{Left: nums[0], Right: nums[1]})
		} else {
			tests = append(tests, nums)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err
	}

	return rules, tests, nil
}

func createLookup(rules []Rule) map[int]Set {
	lookup := make(map[int]Set)
	for _, rule := range rules {
		if val, ok := lookup[rule.Right]; ok {
			val[rule.Left] = struct{}{}
		} else {
			temp := make(Set)
			temp[rule.Left] = struct{}{}
			lookup[rule.Right] = temp
		}
	}
	return lookup
}

func validateTest(lookup map[int]Set, test []int) (bool, int) {
	before := make(Set)
	for idx, elem := range test {
		if toAddList, ok := lookup[elem]; ok {
			for toAdd := range toAddList {
				before[toAdd] = struct{}{}
			}
		}
		if _, ok := before[elem]; ok {
			return false, idx
		}
	}
	return true, -1
}

type Set map[int]struct{}

func part1(filename string) int {
	rules, tests, err := readInput(filename)
	if err != nil {
		return -1
	}
	lookup := createLookup(rules)
	sum := 0
	for _, test := range tests {
		ok, _ := validateTest(lookup, test)
		if ok {
			sum += test[len(test)/2]
		}
	}
	return sum
}

func part2(filename string) int {
	rules, tests, err := readInput(filename)
	if err != nil {
		return -1
	}
	lookup := createLookup(rules)
	sum := 0
	for _, test := range tests {
		ok, ndx := validateTest(lookup, test)
		hadIssue := !ok
		for !ok {
			// 0, 75,97,47,61,53
			// :ndx-1		ndx	ndx-1	ndx+1:
			// 0			97	75		47,61,53
			buffer := []int{}
			buffer = append(buffer, test[:ndx-1]...)
			buffer = append(buffer, test[ndx])
			buffer = append(buffer, test[ndx-1])
			if ndx+1 < len(test) {
				buffer = append(buffer, test[ndx+1:]...)
			}
			test = buffer
			ok, ndx = validateTest(lookup, test)
		}
		if hadIssue {
			sum += test[len(test)/2]
		}
	}
	return sum
}

func main() {
	example := part2("./input1.txt")
	if example != 123 {
		fmt.Println("assertion failed :(", example)
		return
	}
	fmt.Println("assertion passed!!")
	output := part2("input2.txt")
	fmt.Println("output", output)
}
