package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type point struct {
	x float64
	y float64
	z float64
}

type Line struct {
	initial point
	slope   point
}

var (
	lines    = []Line{}
	minSlope float64
	maxSlope float64
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	areaMin := float64(200000000000000)
	areaMax := float64(400000000000000)
	numIntersections := 0
	for i := range lines {
		for j := range lines {
			if i >= j {
				continue
			}
			intersection, t, s := intersectXY(lines[i], lines[j])
			inFuture := t > 0 && s > 0
			if inFuture && intersection.x >= areaMin && intersection.y >= areaMin && intersection.x <= areaMax && intersection.y <= areaMax {
				numIntersections++
			}
		}
	}
	fmt.Printf("%v\n", numIntersections)
}

func part2() {
	// Some "index-finger" reasoning (imagine the line going through all of your fingernails when they are pointed in different directions)
	// A line that intersects 3 other non-parralel lines in 3-D is uniquely determined
	// We have unknown line: L0 := (x0,y0,z0)+m(a0,b0,c0) intersecting L1, L2 and L3 at times r, s, t respectively
	// This gives 9 unknowns but 9 **non-linear** equations (because L0 slope is unknown), which is hard to solve (I tried)
	// Instead, we can brute force over fixed velocities (a0, b0, c0) to get to only 6 unknowns.
	// This is still a lot of velocities to test so we reduce down to only 2 dimensions (X,Y) and consider only intersections with L1 and L2
	// We solve for positions (x0,y0) and times (r,s) using system of linear equations (for each guessed a0, b0)
	// We take those that intersect at positive times r,s and then test those against all other lines to ensure
	// they intersect everywhere in (x,y,z)

	l1 := lines[0]
	x1 := l1.initial.x
	y1 := l1.initial.y
	z1 := l1.initial.z
	a1 := l1.slope.x
	b1 := l1.slope.y
	c1 := l1.slope.z

	l2 := lines[1]
	x2 := l2.initial.x
	y2 := l2.initial.y
	z2 := l2.initial.z
	a2 := l2.slope.x
	b2 := l2.slope.y
	c2 := l2.slope.z

	candidates := []Line{}
	if minSlope < 0 && -minSlope > maxSlope {
		maxSlope = -minSlope
	}
	if maxSlope > 0 && -maxSlope < minSlope {
		minSlope = -maxSlope
	}

	for a0 := int64(minSlope); a0 <= int64(maxSlope); a0++ {
		if a0 == 0 {
			continue
		}
		for b0 := int64(minSlope); b0 <= int64(maxSlope); b0++ {
			if b0 == 0 {
				continue
			}
			v, err := getXYIntersection(float64(a0), float64(b0), x1, y1, a1, b1, x2, y2, a2, b2)
			if err == nil && v[2] > 0 && v[3] > 0 {
				v[0] = math.Round(v[0])
				v[1] = math.Round(v[1])
				l0 := Line{
					initial: point{x: v[0], y: v[1], z: 0},
					slope:   point{x: float64(a0), y: float64(b0)},
				}
				r := math.Round(v[2])
				s := math.Round(v[3])
				// unknowns z0, c0
				// z0 + r*c0 = z1 + r*c1
				// => c0 = (z1+rc1-z0)/r
				// z0 + s*c0 = z2 + s*c2
				// => z0 + s(z1+rc1-z0)/r = z2 + sc2
				// => z0 + sz1/r + src1/r - sz0/r = z2 + sc2
				// => z0(1-s/r)= z2+sc2-sz1/r-sc1
				z0 := math.Round((z2 + s*c2 - s*z1/r - s*c1) / (1 - s/r))
				c0 := math.Round((z1 + r*c1 - z0) / r)
				l0.initial.z = math.Round(z0)
				l0.slope.z = math.Round(c0)
				candidates = append(candidates, l0)
			}
		}
	}

	for _, c := range candidates {
		isValid := true
		for _, line := range lines {
			_, t, s := intersectXY(c, line)
			t = math.Round(t)
			s = math.Round(s)
			if t <= 0 || s <= 0 {
				isValid = false
				break
			}
			x0 := math.Round(c.initial.x + c.slope.x*t)
			x1 := math.Round(line.initial.x + line.slope.x*s)
			if x0 != x1 {
				isValid = false
				break
			}
			y0 := math.Round(c.initial.y + c.slope.y*t)
			y1 := math.Round(line.initial.y + line.slope.y*s)
			if y0 != y1 {
				isValid = false
				break
			}
			z0 := math.Round(c.initial.z + c.slope.z*t)
			z1 := math.Round(line.initial.z + line.slope.z*s)
			if z0 != z1 {
				isValid = false
				break
			}
		}
		if isValid {
			// fmt.Printf("x0:%f, y0: %f, z0: %f, a0: %f, b0: %f, c0: %f\n", c.initial.x, c.initial.y, c.initial.z, c.slope.x, c.slope.y, c.slope.z)
			fmt.Println(c.sum())
		}
	}
}

func (l *Line) sum() int64 {
	return int64(l.initial.x + l.initial.y + l.initial.z)
}

func getXYIntersection(a0, b0, x1, y1, a1, b1, x2, y2, a2, b2 float64) ([]float64, error) {
	// unknowns v=(x0,y0,r,s)
	// x0+a0r=x1+a1r
	// y0+b0r=y1+b1r
	// x0+a0s=x2+a2s
	// y0+b0s=y2+b2s
	// becomes:
	// x0+(a0-a1)r=x1
	// y0+(b0-b1)r=y1
	// x0+(a0-a2)s=x2
	// y0+(b0-b2)s=y2
	// becomes solve Av=b
	// A = [
	//  1,0,a0-a1,0
	//  0,1,b0-b1,0
	//  1,0,0,a0-a2
	//  0,1,0,b0-b2
	// ]
	// b = [x1,y1,x2,y2]
	data := []float64{
		1, 0, a0 - a1, 0,
		0, 1, b0 - b1, 0,
		1, 0, 0, a0 - a2,
		0, 1, 0, b0 - b2,
	}
	A := mat.NewDense(4, 4, data)
	b := mat.NewVecDense(4, []float64{x1, y1, x2, y2})
	v := new(mat.VecDense)
	err := v.SolveVec(A, b)
	if err != nil {
		return nil, err
	}
	return v.RawVector().Data, nil
}

func intersectXY(line1, line2 Line) (pt point, t, s float64) {
	// initial: (x_n, y_n) slope: (a_n, b_n); n=1,2
	// assumption: a_n, b_n != 0
	// line1: x = x_1 + a_1t, y = y_1 + b_1t
	// line2: x = x_2 + a_2s, y = y_2 + b_2s
	// intersection => equal (x,y)
	// => x_1 + a_1t = x_2 + a_2s => s = (x1-x2)/a_2 + (a_1)/(a_2)t
	// => y_1 + b_1t = y_2 + b_2s => s = (y1-y2)/b_2 + (b_1)/(b_2)t
	// => (x1-x2)/a_2 + (a_1)/(a_2)t = (y1-y2)/b_2 + (b_1)/(b_2)t
	// => t = ((x1-x2)/a_2 - (y1-y2)/b_2) / (b1/b2 - a1/a2)

	x1 := line1.initial.x
	y1 := line1.initial.y
	a1 := line1.slope.x
	b1 := line1.slope.y

	x2 := line2.initial.x
	y2 := line2.initial.y
	a2 := line2.slope.x
	b2 := line2.slope.y

	if a1 == 0 || a2 == 0 || b1 == 0 || b2 == 0 {
		panic("did not account for 0 slopes")
	}

	t = (((x1 - x2) / a2) - ((y1 - y2) / b2)) / (b1/b2 - a1/a2)
	s = ((x1 - x2) / a2) + (a1/a2)*t
	pt.x = x1 + a1*t
	pt.y = y1 + b1*t
	return
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " @ ")
		if len(parts) != 2 {
			panic(line)
		}
		l := Line{
			initial: parsePoint(parts[0]),
			slope:   parsePoint(parts[1]),
		}
		lines = append(lines, l)
	}

	for _, l := range lines {
		if l.slope.x < minSlope {
			minSlope = l.slope.x
		}
		if l.slope.x > maxSlope {
			maxSlope = l.slope.x
		}
		if l.slope.y < minSlope {
			minSlope = l.slope.y
		}
		if l.slope.y > maxSlope {
			maxSlope = l.slope.y
		}
		if l.slope.z < minSlope {
			minSlope = l.slope.z
		}
		if l.slope.z > maxSlope {
			maxSlope = l.slope.z
		}
	}
}

func parsePoint(s string) point {
	pt := point{}
	parts := strings.Split(s, ", ")
	for i, p := range parts {
		v, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			panic(err)
		}
		if i == 0 {
			pt.x = float64(v)
		} else if i == 1 {
			pt.y = float64(v)
		} else if i == 2 {
			pt.z = float64(v)
		} else {
			panic(parts)
		}
	}
	return pt
}
