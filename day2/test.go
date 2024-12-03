package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func helper(splits []string, skipped int, n int) (int, error) {
	if skipped < 0 {
		// check from up to 2 elems before current index
		if n > 2 {
			skipped = n - 2
		} else {
			skipped = 0
		}
	} else {
		// increment each iteration
		skipped += 1
	}
	// terminate when either the list ends, or the element to skip
	// is more than one larger than the violated number
	if skipped < len(splits) && skipped < n+2 {
		return processLine(splits, skipped)
	} else {
		return 0, nil
	}
}

func isAscending(splits []string, skipped int) (bool, error) {
	num1, err1 := strconv.Atoi(splits[1])
	num2, err2 := strconv.Atoi(splits[2])
	num0, err3 := strconv.Atoi(splits[0])
	if err1 != nil {
		return false, err1
	}
	if err2 != nil {
		return false, err2
	}
	if err3 != nil {
		return false, err3
	}
	var asc bool
	if skipped == 0 {
		asc = num1 < num2
	} else if skipped == 1 {
		asc = num0 < num2
	} else {
		asc = num0 < num1
	}
	return asc, nil
}

func processLine(splits []string, skipped int) (int, error) {
	prev := -1
	asc, err := isAscending(splits, skipped)
	if err != nil {
		return 0, err
	}

	for n, elem := range splits {
		if n == skipped {
			continue
		}
		i, err := strconv.Atoi(elem)
		if err != nil {
			return 0, err
		}
		if prev >= 0 {
			diff := i - prev
			if !asc {
				diff = -diff
			}
			if diff <= 0 || diff > 3 {
				return helper(splits, skipped, n)
			}
		}
		prev = i
	}
	return 1, nil
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")
		chunk, err := processLine(splits, -1)
		if err != nil {
			fmt.Println("Error parsing int:", err)
			return
		}
		counter += chunk
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Output", strconv.Itoa(counter))
}
