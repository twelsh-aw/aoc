package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	row int
	col int
}

var (
	coords           = map[coord]string{}
	start            = coord{}
	numRows, numCols int
)

type move struct {
	coord coord
	dir   direction
}

type direction string

const (
	north direction = "N"
	south direction = "S"
	east  direction = "E"
	west  direction = "W"
)

type boardResult struct {
	numFilledByParity map[int]int
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	n := 64
	br := solveBoard(n, false)
	total := br.numFilledByParity[n%2]
	fmt.Printf("%v\n", total)
}

func part2() {
	// observations:
	// we start in the middle of the board S := (65,65)
	// the board is a square with odd length X := 131
	// the middle row and middle columns do not contain any walls
	// this means that:
	// - after 0 steps we have 1 empty board
	// - after X/2 steps we have 1 ~filled board and 4 new empty boards (N,E,S,W)
	// - after X/2+X steps we have 5 ~filled boards and 4 half filled boards (NE, NW, SE, SW)
	// - after X/2 + 2X steps we have 13 ~filled boards and 8 half filled boards (NNE, NNW, SSE, SSE, EEN, EES, WWN, WWS)
	// - after X/2 + 3X steps we have 25 ~filled boards and 12 half filled boards
	// - after X/2 + 4X steps we have 41 ~filled boards and 16 half filled boards
	// the number of full boards grows uniformly around the origin in a "diamond" shape, increasing it's perimeter each time
	// moreover, growing this way means we are growing quadratically (kind of like the triangular numbers but in a diamond):
	//   - at (X/2, 3X/2, 5X/2, 7X/2, 9/2X) we have full boards (1, 5, 13, 25, 41) = 1st diffs (4, 8, 12, 16) = 2nd diffs (4, 4, 4) = quadratic
	//   - letting G(N) be number of ~full boards after N steps
	//   - then defining F(n):=G((2n+1)X/2) so that F(0)=G(X/2),F(1)=G(3X/2),F(2)=G(5X/2) we get:
	//   - F(n)=(n+1)^2+n^2
	//   - note this is equivalent to writing: G(N):=F((N-X/2)/X)
	// a few more things:
	//   - since length is odd, every new board starts at an alternating parity; i.e if S was even parity on board X, then on adjacent boards it will be odd parity
	//   - since boards are repeated, after the first initial X/2 steps, the step parities on subsequent boards match the step parities on previous boards (up to alternating behaviour ^)
	//   - the value of N=26501365 is evenly divisible into X after subtracting X/2 i.e the start of our quadratic F
	// we observe that NOT ONLY do the number of ~full boards grow quadratically, but the number of visited points at given parity P does too!
	//   - why exactly?... if you squint, the above observations kind of make this make sense, but squint harder and it's not really clear... some special properties seem to exist beyond the observations above
	//   - but empirically:
	//     - letting g(N) be number of visited gardens at parity P after N steps
	//     - letting f(n) be number of visited gardens at parity P after n "groups of steps of size X",
	//   - then we have for parity 1:
	//     - f(0), f(1), f(2), f(3), f(4) = g(X/2), g(3X/2), g(5X/2), g(7X/2), g(9X/2)
	//     = 3778, 33695, 93438, 183007, 302402
	//     1st diffs: = 29917, 59743, 89569, 119395
	//     2nd diffs: = 29826, 29826, 29826 ==> quadratic
	N := 26501365
	X := numRows
	n := (N - X/2) / X
	p1 := solveBoard(X/2, true)
	p2 := solveBoard(3*X/2, true)
	p3 := solveBoard(5*X/2, true)
	// p4 := solveBoard(7*X/2, true)
	// p5 := solveBoard(9*X/2, true)
	points := [3]int{
		p1.numFilledByParity[N%2],
		p2.numFilledByParity[(N+1)%2],
		p3.numFilledByParity[N%2],
	}
	numFilled := fitQuadraticFromPoints(n, points)
	fmt.Println(numFilled)
}

func solveBoard(n int, allowInfinite bool) boardResult {
	seen := map[coord]bool{
		start: true,
	}
	cur := []move{
		{start, ""},
	}
	numFilledByParity := map[int]int{
		0: 1,
	}
	br := boardResult{
		numFilledByParity: numFilledByParity,
	}
	for i := 1; i <= n; i++ {
		next := []move{}
		for _, m := range cur {
			nm := m.getNextValidMoves(seen, allowInfinite)
			next = append(next, nm...)
		}
		if len(next) == 0 {
			return br
		}
		cur = append([]move{}, next...)
		numFilledByParity[i%2] += len(next)
	}
	return br
}

func (m *move) getNextValidMoves(seen map[coord]bool, allowInfinite bool) []move {
	moves := []move{}
	dirs := []direction{north, south, east, west}
	for _, dir := range dirs {
		// if we previously got to coord by making move in dir, we don't need to go backwards
		if m.dir == dir.opposite() {
			continue
		}
		next := m.coord
		switch dir {
		case north:
			next.row--
		case south:
			next.row++
		case east:
			next.col++
		case west:
			next.col--
		}
		val := coords[next]
		if val == "" && allowInfinite {
			val = coords[next.getMod()]
		} else if val == "" {
			continue
		}

		if seen[next] || val == "#" {
			continue
		}
		moves = append(moves, move{next, dir})
		seen[next] = true
	}
	return moves
}

func (c *coord) getMod() coord {
	mod := coord{c.row % numRows, c.col % numCols}
	if mod.row < 0 {
		mod.row += numRows
	}
	if mod.col < 0 {
		mod.col += numCols
	}
	return mod
}

func (d direction) opposite() direction {
	switch d {
	case north:
		return south
	case east:
		return west
	case west:
		return east
	case south:
		return north
	default:
		panic(d)
	}
}

// see https://en.wikipedia.org/wiki/Newton_polynomial
func fitQuadraticFromPoints(x int, a [3]int) int {
	b0 := a[0]
	b1 := a[1] - a[0]
	b2 := a[2] - a[1]

	return b0 + (b1 * x) + (x*(x-1)/2)*(b2-b1)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for row, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		numRows++
		parts := strings.Split(line, "")
		numCols = len(parts)
		for col, v := range parts {
			co := coord{row, col}
			coords[co] = v
			if v == "S" {
				start = co
			}
		}
	}
}
