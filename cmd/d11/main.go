package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jacobmiller22/aoc-2024/numeric"
)

func main() {
	path := flag.String("input", "", "Path to the input date")

	flag.Parse()

	if *path == "" {
		log.Fatalf("Input path must be non-empty")
	}

	f, err := os.Open(*path)
	if err != nil {
		log.Fatalf("error opening file: %v", f)
	}

	p1, p2, err := solution(f)
	if err != nil {
		log.Fatalf("error running solution: %v", err)
	}

	log.Printf("Results: p1=%d; p2=%d\n", p1, p2)
}

func parseRocks(f io.Reader) ([]int, error) {
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	line := strings.Trim(string(data), "\n")
	strRocks := strings.Split(line, " ")

	rocks := make([]int, 0, len(strRocks))

	for _, s := range strRocks {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("error parsing rocks input: %v", err)
		}
		rocks = append(rocks, n)
	}
	return rocks, nil
}

type RockStep struct {
	rock  int
	steps int
}

func cached(f func(RockStep) int) func(RockStep) int {
	cache := make(map[RockStep]int)
	return func(x RockStep) int {
		if _, ok := cache[x]; !ok {
			cache[x] = f(x)
		}
		value, _ := cache[x]
		return value
	}
}

var (
	blink       func(RockStep) int
	blinkCached func(RockStep) int
)

func solution(f io.Reader) (int, int, error) {
	rocks, err := parseRocks(f)
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing rocks : %v", err)
	}

	blink := func(rs RockStep) int {
		if rs.steps == 0 {
			return 1
		}
		if rs.rock == 0 {
			return blinkCached(RockStep{1, rs.steps - 1})
		}
		ndigits := numeric.NDigits(float64(rs.rock))
		if ndigits%2 == 0 {
			l, r := numeric.SplitNum(rs.rock)
			return blinkCached(RockStep{l, rs.steps - 1}) + blinkCached(RockStep{r, rs.steps - 1})

		}
		return blinkCached(RockStep{rs.rock * 2024, rs.steps - 1})
	}
	blinkCached = cached(blink)

	sum1 := 0
	sum2 := 0
	for _, r := range rocks {
		sum1 += blinkCached(RockStep{r, 25})
		sum2 += blinkCached(RockStep{r, 75})
	}

	return sum1, sum2, nil
}
