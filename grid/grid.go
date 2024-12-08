package grid

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) X() int {
	return c.x
}

func (c *Coordinate) Y() int {
	return c.y
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{x, y}
}

type Grid struct {
	m map[Coordinate]bool
	h int
	w int
}

func NewGrid() *Grid {
	return &Grid{
		m: make(map[Coordinate]bool),
	}
}

func (b *Grid) Clear() {
	for k := range b.m {
		b.m[k] = false
	}
}

func (b *Grid) Mark(x, y int) {
	coord := Coordinate{x, y}
	b.m[coord] = true
}

func (b *Grid) Unmark(x, y int) {
	coord := Coordinate{x, y}
	b.m[coord] = false
}

func (b *Grid) Has(x, y int) bool {
	coord := Coordinate{x, y}
	v, _ := b.m[coord]
	return v
}

func generateCombinations[T comparable, V any](m map[T]V, n int) [][]T {
	keys := make([]T, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	var result [][]T
	var comb func(start int, current []T)

	comb = func(start int, current []T) {
		if len(current) == n {
			combination := make([]T, n)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		for i := start; i < len(keys); i++ {
			comb(i+1, append(current, keys[i]))
		}
	}

	comb(0, []T{})
	return result
}

func (b *Grid) Len() int {
	return len(b.m)
}

func (b *Grid) Pairs() [][]Coordinate {
	return generateCombinations(b.m, 2)
}

func (b *Grid) Height() int {
	return b.h
}

func (b *Grid) SetHeight(h int) {
	b.h = h
}

func (b *Grid) Width() int {
	return b.w
}
func (b *Grid) SetWidth(w int) {
	b.w = w
}

type ClassifiedGrid struct {
	m map[rune]*Grid
	h int
	w int
}

func (b *ClassifiedGrid) Height() int {
	return b.h
}

func (b *ClassifiedGrid) SetHeight(h int) {
	b.h = h
}

func (b *ClassifiedGrid) Width() int {
	return b.w
}
func (b *ClassifiedGrid) SetWidth(w int) {
	b.w = w
}

func (g *ClassifiedGrid) Clear() {
	for _, v := range g.m {
		v.Clear()
	}
}

func (g *ClassifiedGrid) Mark(c rune, x, y int) {
	if _, ok := g.m[c]; !ok {
		g.m[c] = NewGrid()
	}
	g.m[c].Mark(x, y)
}

func (g *ClassifiedGrid) Unmark(c rune, x, y int) {
	if m, ok := g.m[c]; ok {
		m.Unmark(x, y)
	}
}

func (g *ClassifiedGrid) Has(c rune, x, y int) bool {
	if m, ok := g.m[c]; ok {
		return m.Has(x, y)
	}
	return false
}

func (g *ClassifiedGrid) Pairs(c rune) [][]Coordinate {
	if m, ok := g.m[c]; ok {
		return m.Pairs()
	}
	return nil
}

func (g *ClassifiedGrid) Grid() map[rune]*Grid {
	return g.m
}

func NewClassifiedGrid() *ClassifiedGrid {
	return &ClassifiedGrid{m: make(map[rune]*Grid, 0)}
}
