package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	path := flag.String("input", "", "Path to input")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}

	part1, part2, err := problem(f)

	if err != nil {
		log.Fatalf("error running part: %v", err)
	}

	log.Printf("Results:\n\tPart 1: %d\n\tPart 2: %d\n", part1, part2)

}

// finds `target` string in string slice `s`
// returns true if found
func find(s []string, target string) bool {
	for _, e := range s {
		if e == target {
			return true
		}
	}
	return false
}

func isValid(deps map[string][]string, elems []string) bool {
	for i, elem := range elems {
		// for each elem,
		// check if any dep is before this, its invalid
		for _, dep := range deps[elem] {
			if find(elems[:i], dep) {
				// invalid
				return false
			}
		}
	}
	return true
}

func middle(s []string) (int, error) {
	middle, err := strconv.Atoi(s[len(s)/2])
	if err != nil {
		return 0, err
	}
	return middle, err

}

func problem(f io.Reader) (int, int, error) {

	deps := make(map[string][]string, 0)

	scr := bufio.NewScanner(f)

	for scr.Scan() {
		line := scr.Text()

		if line == "" {
			break
		}

		splitStrs := strings.Split(line, "|")
		left := splitStrs[0]
		right := splitStrs[1]
		// left|right => left must come before right => left depends on right

		if _, ok := deps[left]; !ok {
			deps[left] = make([]string, 0, 1)
		}
		deps[left] = append(deps[left], right)
	}

	if err := scr.Err(); err != nil {
		return 0, 0, err
	}

	// map is created
	sum1 := 0
	sum2 := 0

	for scr.Scan() {
		line := scr.Text()

		if line == "" {
			break
		}

		elems := strings.Split(line, ",")

		if isValid(deps, elems) {
			mid, err := middle(elems)
			if err != nil {
				return 0, 0, err
			}
			sum1 += mid

		} else {
			// part 2
		Outer:
			for i := 0; i < len(elems)-1; i++ {
				for j := 0; j < len(elems)-i-1; j++ {
					if dep, ok := deps[elems[j+1]]; ok {
						if find(dep, elems[j]) {
							temp := elems[j]
							elems[j] = elems[j+1]
							elems[j+1] = temp
						}
						if isValid(deps, elems) {
							mid, err := middle(elems)
							if err != nil {
								return 0, 0, err
							}
							sum2 += mid
							break Outer
						}

					}
				}
			}

		}

	}

	if err := scr.Err(); err != nil {
		return 0, 0, err
	}

	return sum1, sum2, nil
}
