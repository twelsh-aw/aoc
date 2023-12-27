package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type graph struct {
	V map[string]bool
	E map[string]map[string]bool
}

type edge struct {
	from string
	to   string
}

var (
	g     = graph{V: map[string]bool{}, E: map[string]map[string]bool{}}
	edges = []edge{}
	start = ""
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	source := []string{start}
	sink := []string{}
	// we have a 3-connected graph
	// we can partition the verticies into those that are 3-connected to source and those that are 4+-connected to source
	// these are the same groups that would remain after a 3-cut separating source from some sink on the other side
	// once we know the groups, we can multiply them. I don't actually find the 3-cut

	// let G1, G2 be the sizes of connected subgraphs after the min cut is made such that G1*G2 is unique.
	// Pick source vertex s such that s belongs to G1 and G1<=G2.
	// let v != s be another vertex and let n be the number of (s,v)-edge-disjoint paths.
	// then n=3 vs n>=4 partititions V into sets of size G1,G2.
	// proof:
	// suppose n >= 4 and v in G2. Then after the min cut, G1 and G2 are still connected, a contradiction.
	// suppose n = 3 and v in G1. Then we can find a 3-edge (Mengers) cut in G1 separating s and v.
	// this gives new sizes of sets (G1-x), (G2+x).
	// Thus (G1-x)(G2+x)=G1G2 by uniqueness.
	// => (G1-G2)x-x^2=0 => x(x-(G1-G2))=0
	// => x = 0 OR x=G1-G2 => x <= 0, which contradicts that v is in G1.
	for v := range g.V {
		if v == start {
			continue
		}
		n := g.findNumEdgeDisjointPaths(start, v)
		if n == 3 {
			sink = append(sink, v)
		} else if n == 4 {
			source = append(source, v)
		} else {
			panic(v)
		}
	}

	n := len(source) * len(sink)
	fmt.Println(n)
}

func part2() {
	fmt.Printf("%v\n", nil)
}

func (g *graph) findNumEdgeDisjointPaths(source, dest string) int {
	// by Mengers, the number of edge disjoint paths is equal to number of edges separating source from dest
	// just iteratively apply Djikstra and remove the shortest path each time to get number of edge disjoint paths.
	type pos struct {
		prev    map[string]string
		dists   map[string]int
		visited map[string]bool
	}
	toRemove := []edge{}
	numPaths := 0
	for i := 0; i < 4; i++ {
		gt := g.clone()
		for _, e := range toRemove {
			delete(gt.E[e.from], e.to)
			delete(gt.E[e.to], e.from)
		}
		cur := pos{
			visited: map[string]bool{},
			prev:    map[string]string{},
			dists: map[string]int{
				start: 0,
			},
		}
		for {
			var v string
			minDist := math.MaxInt
			for vertex, dist := range cur.dists {
				if dist < minDist {
					minDist = dist
					v = vertex
				}
			}
			if v == dest {
				to := dest
				for to != source {
					e := edge{}
					e.to = to
					e.from = cur.prev[to]
					toRemove = append(toRemove, e)
					to = e.from
				}
				numPaths++
				break
			}
			if v == "" {
				// no more paths
				break
			}

			cur.visited[v] = true
			delete(cur.dists, v)
			for a := range gt.E[v] {
				if !cur.visited[a] {
					dist, ok := cur.dists[a]
					if !ok {
						dist = math.MaxInt
					}
					if minDist+1 < dist {
						cur.dists[a] = minDist + 1
						cur.prev[a] = v
					}
				}
			}
		}
	}
	return numPaths
}

func (g *graph) clone() graph {
	cloned := graph{V: map[string]bool{}, E: map[string]map[string]bool{}}
	for k := range g.V {
		cloned.V[k] = true
	}
	for k := range g.E {
		cloned.E[k] = map[string]bool{}
		for e := range g.E[k] {
			cloned.E[k][e] = true
		}
	}
	return cloned
}
func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	eMap := map[string]bool{}
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		p := strings.Split(line, ": ")
		v0 := p[0]
		if !g.V[v0] {
			g.V[v0] = true
			g.E[v0] = map[string]bool{}
		}
		if start == "" {
			start = v0
		}

		e := strings.Split(p[1], " ")
		for _, v := range e {
			if !g.V[v] {
				g.V[v] = true
				g.E[v] = map[string]bool{}
			}
			g.E[v][v0] = true
			g.E[v0][v] = true
			from := v0
			to := v
			if v < v0 {
				from = v
				to = v0
			}
			ed := fmt.Sprintf("%s:%s", from, to)
			if !eMap[ed] {
				eMap[ed] = true
				edges = append(edges, edge{from: from, to: to})
			}
		}
	}
}
