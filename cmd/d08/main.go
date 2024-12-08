package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jacobmiller22/aoc-2024/grid"
	"github.com/jacobmiller22/aoc-2024/math"
)

func main() {
	path := flag.String("input", "", "path to input")

	flag.Parse()

	if *path == "" {
		log.Fatalf("provided input empty")
	}

	f, err := os.Open(*path)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	r1, r2, err := solution(f)

	if err != nil {
		log.Fatalf("error running solution: %v", err)
	}

	fmt.Printf("Results: Part 1=%d; Part 2=%d\n", r1, r2)
}

func scanAnts(f io.Reader) (*grid.ClassifiedGrid, error) {

	scr := bufio.NewScanner(f)

	g := grid.NewClassifiedGrid()

	y := 0
	maxX := 0
	for scr.Scan() {
		line := scr.Text()

		for x, c := range line {
			maxX = math.Max(maxX, x)

			if c == '.' {
				continue
			}
			g.Mark(c, x, y)
		}

		y++
	}

	if err := scr.Err(); err != nil {
		return nil, fmt.Errorf("error scanning input: %v", err)
	}
	g.SetWidth(maxX + 1)
	g.SetHeight(y)

	return g, nil
}

func visualizeAntinodes(anG *grid.Grid, fg *grid.ClassifiedGrid) {

	fmt.Println("\n---Board Viewer---")
	for col := 0; col < fg.Width(); col++ {
		line := ""
		for row := 0; row < fg.Height(); row++ {
			if anG.Has(row, col) {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}

}

func solution(f io.Reader) (int, int, error) {
	fg, err := scanAnts(f)
	if err != nil {
		return 0, 0, fmt.Errorf("error while parsing input: %v", err)
	}

	// for each antennae, create another board for each unique antinode, use the same board for each frequency
	anG := grid.NewGrid()

	for _, v := range fg.Grid() {
		pairs := v.Pairs()
		for _, p := range pairs {
			dx, dy := p[0].X()-p[1].X(), p[0].Y()-p[1].Y()
			x1, y1 := p[0].X()+dx, p[0].Y()+dy
			if x1 >= 0 && x1 < fg.Width() && y1 >= 0 && y1 < fg.Height() {
				anG.Mark(x1, y1)

			}
			x2, y2 := p[1].X()-dx, p[1].Y()-dy
			if x2 >= 0 && x2 < fg.Width() && y2 >= 0 && y2 < fg.Height() {
				anG.Mark(x2, y2)
			}
		}
	}
	sum1 := anG.Len()

	for _, v := range fg.Grid() {
		pairs := v.Pairs()
		for _, p := range pairs {
			for i := 1; ; i++ {
				dx, dy := (p[0].X()-p[1].X())*i, (p[0].Y()-p[1].Y())*i
				x1, y1 := p[0].X()+dx, p[0].Y()+dy
				x2, y2 := p[1].X()-dx, p[1].Y()-dy

				firstInBounds := x1 >= 0 && x1 < fg.Width() && y1 >= 0 && y1 < fg.Height()
				if firstInBounds {
					anG.Mark(x1, y1)
				}
				secondInBounds := x2 >= 0 && x2 < fg.Width() && y2 >= 0 && y2 < fg.Height()
				if secondInBounds {
					anG.Mark(x2, y2)
				}
				if !firstInBounds && !secondInBounds {
					break
				}
			}
			anG.Mark(p[0].X(), p[0].Y())
			anG.Mark(p[1].X(), p[1].Y())

		}
	}

	return sum1, anG.Len(), nil
}
