package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/jacobmiller22/aoc-2024/grid"
)

func main() {
	path := flag.String("input", "", "Path to the input file")

	flag.Parse()

	if *path == "" {
		log.Fatal("Provided input cannot be empty.\n")
	}

	f, err := os.Open(*path)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	p1, p2, err := solution(f)
	if err != nil {
		log.Fatalf("error running solution: %v", err)
	}

	log.Printf("Results: p1=%d; p2=%d\n", p1, p2)
}

const (
	InfoColor    = "\033[1;34m"
	NoticeColor  = "\033[1;36m"
	WarningColor = "\033[1;33m"
	ErrorColor   = "\033[1;31m"
	DebugColor   = "\033[0;36m"
	ResetColor   = "\033[0m"
)

func visualize(M *[]string, x, y int, t byte, visited *grid.Grid) {
	var w strings.Builder
	w.WriteString("------------------\n-- Visualizer --\n------------------\n")
	for r := 0; r < len(*M); r++ {
		for c := 0; c < len((*M)[0]); c++ {

			if r == x && c == y {
				w.WriteString(WarningColor)

			} else if visited.Has(r, c) {
				w.WriteString(DebugColor)
			} else if (*M)[r][c] == t {
				w.WriteString(ErrorColor)
			}

			w.WriteString(string((*M)[r][c]))
			w.WriteString(ResetColor)
		}
		w.WriteString("\n")
	}
	fmt.Print(w.String())
}

func dfsIterative(M *[]string, x, y int, visited *grid.Grid) int {
	stack := make([]grid.Coordinate, 0, 0)
	stack = append(stack, *grid.NewCoordinate(x, y))

	sum := 0

	for len(stack) > 0 {
		// Pop stack
		s := stack[len(stack)-1]
		stack = slices.Delete(stack, len(stack)-1, len(stack))
		r, c := s.X(), s.Y()
		v := (*M)[r][c] // Value of current node

		// visualize(M, r, c, v, visited)
		if !visited.Has(r, c) {
			if v == '9' {
				sum++
			}
			visited.Mark(r, c)
		}

		// push unvisted adjencent steps to stack
		if r-1 >= 0 && (*M)[r-1][c] == v+1 {
			stack = append(stack, *grid.NewCoordinate(r-1, c))
		}

		// down
		if r < len(*M)-1 && (*M)[r+1][c] == v+1 {
			stack = append(stack, *grid.NewCoordinate(r+1, c))
		}

		// left
		if c-1 >= 0 && (*M)[r][c-1] == v+1 {
			stack = append(stack, *grid.NewCoordinate(r, c-1))
		}

		// right
		if c < len((*M)[0])-1 && (*M)[r][c+1] == v+1 {
			stack = append(stack, *grid.NewCoordinate(r, c+1))
		}
	}
	return sum
}

func dfs(M *[]string, r, c int, t byte, visited *grid.Grid) int {
	// visualize(M, r, c, t, visited)

	//  Check if we have been here before, if we have a visited map
	if visited != nil {
		if visited.Has(r, c) {
			return 0
		}

		visited.Mark(r, c)
	}

	// base case, look for '9'
	if (*M)[r][c] == '9' {
		return 1
	}

	s := 0
	// up
	if r-1 >= 0 && (*M)[r-1][c] == t+1 {
		s += dfs(M, r-1, c, t+1, visited)
	}

	// down
	if r < len(*M)-1 && (*M)[r+1][c] == t+1 {
		s += dfs(M, r+1, c, t+1, visited)
	}

	// left
	if c-1 >= 0 && (*M)[r][c-1] == t+1 {
		s += dfs(M, r, c-1, t+1, visited)
	}

	// right
	if c < len((*M)[0])-1 && (*M)[r][c+1] == t+1 {
		s += dfs(M, r, c+1, t+1, visited)
	}

	return s

}

func solution(f io.Reader) (int, int, error) {
	scr := bufio.NewScanner(f)

	M := make([]string, 0, 0)
	for scr.Scan() {
		line, _ := strings.CutSuffix(scr.Text(), "\n")
		M = append(M, line)
	}

	if err := scr.Err(); err != nil {
		log.Fatalf("error scanning file: %v", err)
	}

	sum1 := 0
	sum2 := 0
	for r := 0; r < len(M); r++ {
		for c := 0; c < len(M[0]); c++ {
			// Find the trail heads
			if M[r][c] == '0' {
				// found trailhead
				visited := grid.NewGrid()
				visited.SetWidth(len(M[0]))
				visited.SetHeight(len(M))
				// sum1 += dfs(&M, r, c, '0', visited)
				sum1 += dfsIterative(&M, r, c, visited)
				sum2 += dfs(&M, r, c, '0', nil)

			}
		}
	}

	return sum1, sum2, nil
}
