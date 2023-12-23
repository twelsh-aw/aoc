package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

var (
	coords           = map[coord]string{}
	numRows, numCols int
	start, end       coord
)

type coord struct {
	row int
	col int
}

type pos struct {
	coord   coord
	steps   int
	visited map[coord]bool
}

type graph struct {
	verticies   map[coord]bool
	edgeWeights map[coord]map[coord]int
}

type coordHistory struct {
	cur  coord
	prev coord
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	g := makeGraph(true)
	n := g.getLongestPath()
	fmt.Printf("%v\n", n)
}

func part2() {
	g := makeGraph(false)
	n := g.getLongestPath()
	fmt.Printf("%v\n", n)
}

// makeGraph creates a new graph where:
// the verticies are coords where a split can be made
// an edge exists between u,v iff there is a (single) path from u to v in the coords grid
// the weight of the edge is the number of steps in coords grid for that path
// this graph reduction greatly simplifies the number of paths we have to look at, making part2 feasible
func makeGraph(withSlopes bool) graph {
	g := graph{
		verticies: map[coord]bool{
			start: true,
			end:   false,
		},
		edgeWeights: map[coord]map[coord]int{
			start: make(map[coord]int, 0),
			end:   make(map[coord]int, 0),
		},
	}
	verticesToReduce := []coordHistory{
		{
			cur:  coord{1, 1},
			prev: start,
		},
	}
	seen := map[coordHistory]bool{}
	for len(verticesToReduce) > 0 {
		nextVerticies := []coordHistory{}
		for _, vertex := range verticesToReduce {
			seen[vertex] = true
			v := vertex.prev
			cur := pos{
				coord: vertex.cur,
				steps: 1,
				visited: map[coord]bool{
					vertex.prev: true,
				},
			}
			for {
				if cur.coord == end {
					if cur.steps > g.edgeWeights[v][end] {
						g.edgeWeights[v][end] = cur.steps
					}
					break
				}
				nb := cur.getNeighbours(withSlopes)
				if len(nb) == 0 {
					break
				}
				cur.visited[cur.coord] = true
				if len(nb) == 1 {
					cur.coord = nb[0]
					cur.steps++
					continue
				}
				if cur.steps > g.edgeWeights[v][cur.coord] {
					g.edgeWeights[v][cur.coord] = cur.steps
				}
				if !g.verticies[cur.coord] {
					g.verticies[cur.coord] = true
					g.edgeWeights[cur.coord] = map[coord]int{}
				}
				for _, n := range nb {
					nv := coordHistory{
						cur:  n,
						prev: cur.coord,
					}
					if !seen[nv] {
						nextVerticies = append(nextVerticies, nv)
					}
				}
				break
			}
		}
		verticesToReduce = append([]coordHistory{}, nextVerticies...)
	}
	return g
}

func (g *graph) getLongestPath() int {
	cur := []pos{
		{
			coord:   start,
			steps:   0,
			visited: map[coord]bool{},
		},
	}
	var max int
	for len(cur) > 0 {
		next := []pos{}
		for _, c := range cur {
			c.visited[c.coord] = true
			if c.coord == end {
				if c.steps > max {
					fmt.Println("new max", c.steps)
					max = c.steps
				}
				continue
			}
			adj := g.edgeWeights[c.coord]
			for n, w := range adj {
				if c.visited[n] {
					continue
				}
				next = append(next, pos{
					coord:   n,
					steps:   c.steps + w,
					visited: maps.Clone(c.visited),
				})
			}
		}
		cur = append([]pos{}, next...)
	}

	return max
}

func (p *pos) getNeighbours(withSlopes bool) []coord {
	c := p.coord
	val := coords[c]
	up, down, left, right := c, c, c, c
	right.col++
	left.col--
	down.row++
	up.row--
	next := []coord{}
	allowAnyDir := !withSlopes || val == "."
	if allowAnyDir || val == ">" {
		valid := right.valid() && (coords[right] == "." || coords[right] == ">" || (!withSlopes && coords[right] != "#"))
		if valid && !p.visited[right] {
			next = append(next, right)
		}
	}
	if allowAnyDir || val == "<" {
		valid := left.valid() && (coords[left] == "." || coords[left] == "<" || (!withSlopes && coords[left] != "#"))
		if valid && !p.visited[left] {
			next = append(next, left)
		}
	}
	if allowAnyDir || val == "^" {
		valid := up.valid() && (coords[up] == "." || coords[up] == "^" || (!withSlopes && coords[up] != "#"))
		if valid && !p.visited[up] {
			next = append(next, up)
		}
	}
	if allowAnyDir || val == "v" {
		valid := down.valid() && (coords[down] == "." || coords[down] == "v" || (!withSlopes && coords[down] != "#"))
		if valid && !p.visited[down] {
			next = append(next, down)
		}
	}
	return next
}

func (c coord) valid() bool {
	return c.row > 0 && c.col > 0 && c.row < numRows && c.col < numCols
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
			c := coord{row, col}
			if row == 0 && v == "." {
				start = c
			} else if v == "." {
				end = c
			}
			coords[c] = v
		}
	}
}
