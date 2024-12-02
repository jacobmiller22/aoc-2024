package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jacobmiller22/aoc-2024/math"
)

func main() {

	path := flag.String("input", "", "The path to the file to run")
	part := flag.Int("part", 1, "The problem part to run")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}

	var result int
	switch *part {
	case 1:
		result, err = part1(f)
	case 2:
		result, err = part2(f)
	}

	if err != nil {
		log.Fatalf("Error running part: %v", err)
	}

	log.Printf("Result for part %d: %d", *part, result)

}

func isGradual(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if math.Abs(nums[i]-nums[i-1]) > 3 {
			return false
		}
		if math.Abs(nums[i]-nums[i-1]) == 0 {
			return false
		}
	}
	return true
}

func isIncreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func isDecreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			return false
		}
	}
	return true
}

func isValid(nums []int) bool {
	inc, dec, gradual := isIncreasing(nums), isDecreasing(nums), isGradual(nums)
	return (inc || dec) && gradual

}

func part1(f io.Reader) (int, error) {
	scanner := bufio.NewScanner(f)

	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		strNums := strings.Split(line, " ")

		nums := make([]int, 0, 0)
		for _, strInt := range strNums {
			currInt, err := strconv.Atoi(strInt)
			if err != nil {
				return 0, fmt.Errorf("error converting string to number: %v", err)
			}
			nums = append(nums, currInt)
		}

		if isValid(nums) {
			safeCount += 1
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning: %v", err)
	}

	return safeCount, nil
}

func part2(f io.Reader) (int, error) {
	scanner := bufio.NewScanner(f)

	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		strInts := strings.Split(line, " ")

		nums := make([]int, 0, 0)
		for _, strInt := range strInts {
			currInt, err := strconv.Atoi(strInt)
			if err != nil {
				return 0, fmt.Errorf("error converting string to number: %v", err)
			}
			nums = append(nums, currInt)
		}

		for i := 0; i < len(nums); i++ {
			l := append([]int{}, nums[:i]...)
			r := append(l, nums[i+1:]...)
			if isValid(r) {
				safeCount += 1
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning: %v", err)
	}

	return safeCount, nil
}
