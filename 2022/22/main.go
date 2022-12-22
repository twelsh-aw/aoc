package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	row int
	col int
}

type cubeCoord struct {
	coord
	face int
}

type cube struct {
	coords        map[cubeCoord]string
	adjacentFaces map[int]map[direction]adjacentFace
	faceRanges    map[int]faceRange
}

type adjacentFace struct {
	face         int
	newDirection direction
	transform    func(c coord) coord
}

type instruction struct {
	steps    int
	rotation rotation
}

type faceRange struct {
	rowRange coordRange
	colRange coordRange
}

type coordRange struct {
	min int
	max int
}

type direction string
type rotation string

const (
	right direction = "R"
	down  direction = "D"
	left  direction = "L"
	up    direction = "U"

	clockwise        rotation = "R"
	counterClockwise rotation = "L"
	none             rotation = ""
)

var (
	coords        = make(map[coord]string)
	instructions  []instruction
	rowRanges     = make(map[int]coordRange)
	colRanges     = make(map[int]coordRange)
	adjacentFaces = make(map[int]map[direction]adjacentFace)
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	curPosition := coord{1, rowRanges[1].min}
	curDirection := right
	canMove := true
	for _, ins := range instructions {
		curDirection = getNextDirection(curDirection, ins.rotation)
		for i := 0; i < ins.steps; i++ {
			curPosition, canMove = move2D(curPosition, curDirection)
			if !canMove {
				break
			}
		}
	}

	password := (curPosition.row * 1000) + (curPosition.col * 4) + curDirection.asValue()
	fmt.Printf("%v\n", password)
}

func part2() {
	cbe := foldCube(50)
	curPosition := cubeCoord{coord: coord{1, rowRanges[1].min}, face: 1}
	curDirection := right
	nextDirection := right
	canMove := true
	for _, ins := range instructions {
		curDirection = getNextDirection(curDirection, ins.rotation)
		for i := 0; i < ins.steps; i++ {
			curPosition, nextDirection, canMove = move3D(curPosition, curDirection, cbe)
			if !canMove {
				break
			}

			curDirection = nextDirection
		}
	}
	password := (curPosition.row * 1000) + (curPosition.col * 4) + curDirection.asValue()
	fmt.Printf("%v\n", password)
}

func move3D(pos cubeCoord, dir direction, cbe cube) (cubeCoord, direction, bool) {
	currentFace := pos.face
	currentFaceRange := cbe.faceRanges[currentFace]
	currentRowNormalized := pos.row - currentFaceRange.rowRange.min
	currentColNormalized := pos.col - currentFaceRange.colRange.min

	next := pos
	nextDir := dir
	switch dir {
	case right:
		next.col++
		if next.col > currentFaceRange.colRange.max {
			adj := cbe.adjacentFaces[pos.face][right]
			next.face = adj.face
			nextDir = adj.newDirection
			nextFaceRange := cbe.faceRanges[adj.face]
			nextNorm := adj.transform(coord{currentRowNormalized, currentColNormalized})
			next.row = nextFaceRange.rowRange.min + nextNorm.row
			next.col = nextFaceRange.colRange.min + nextNorm.col
		}
	case left:
		next.col--
		if next.col < currentFaceRange.colRange.min {
			adj := cbe.adjacentFaces[pos.face][left]
			next.face = adj.face
			nextDir = adj.newDirection
			nextFaceRange := cbe.faceRanges[adj.face]
			nextNorm := adj.transform(coord{currentRowNormalized, currentColNormalized})
			next.row = nextFaceRange.rowRange.min + nextNorm.row
			next.col = nextFaceRange.colRange.min + nextNorm.col
		}
	case down:
		next.row++
		if next.row > currentFaceRange.rowRange.max {
			adj := cbe.adjacentFaces[pos.face][down]
			next.face = adj.face
			nextDir = adj.newDirection
			nextFaceRange := cbe.faceRanges[adj.face]
			nextNorm := adj.transform(coord{currentRowNormalized, currentColNormalized})
			next.row = nextFaceRange.rowRange.min + nextNorm.row
			next.col = nextFaceRange.colRange.min + nextNorm.col
		}
	case up:
		next.row--
		if next.row < currentFaceRange.rowRange.min {
			adj := cbe.adjacentFaces[pos.face][up]
			next.face = adj.face
			nextDir = adj.newDirection
			nextFaceRange := cbe.faceRanges[adj.face]
			nextNorm := adj.transform(coord{currentRowNormalized, currentColNormalized})
			next.row = nextFaceRange.rowRange.min + nextNorm.row
			next.col = nextFaceRange.colRange.min + nextNorm.col
		}
	}

	tile := cbe.coords[next]
	if tile == "." {
		return next, nextDir, true
	} else if tile == "#" {
		return pos, nextDir, false
	} else {
		panic(pos)
	}
}

func move2D(pos coord, dir direction) (coord, bool) {
	next := pos
	switch dir {
	case right:
		next.col++
		if next.col > rowRanges[pos.row].max {
			next.col = rowRanges[pos.row].min
		}
	case left:
		next.col--
		if next.col < rowRanges[pos.row].min {
			next.col = rowRanges[pos.row].max
		}
	case down:
		next.row++
		if next.row > colRanges[pos.col].max {
			next.row = colRanges[pos.col].min
		}
	case up:
		next.row--
		if next.row < colRanges[pos.col].min {
			next.row = colRanges[pos.col].max
		}
	}

	tile := coords[next]
	if tile == "." {
		return next, true
	} else if tile == "#" {
		return pos, false
	} else {
		panic(pos)
	}
}

func foldCube(faceSize int) cube {
	cbe := cube{
		coords:     make(map[cubeCoord]string),
		faceRanges: make(map[int]faceRange),
	}
	//facesByID := map[int]coord{
	//	1: {0, 2},
	//	2: {1, 0},
	//	3: {1, 1},
	//	4: {1, 2},
	//	5: {2, 2},
	//	6: {2, 3},
	//}
	//
	//cbe.adjacentFaces = map[int]map[direction]adjacentFace{
	//	1: {
	//		down: {
	//			face:         4,
	//			newDirection: down,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: 0,
	//					col: c.col,
	//				}
	//			},
	//		},
	//		left: {
	//			face:         3,
	//			newDirection: down,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: 0,
	//					col: c.row,
	//				}
	//			},
	//		},
	//		up: {
	//			face:         2,
	//			newDirection: down,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: 0,
	//					col: faceSize - c.col - 1,
	//				}
	//			},
	//		},
	//		right: {
	//			face:         6,
	//			newDirection: left,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - c.row - 1,
	//					col: faceSize - 1,
	//				}
	//			},
	//		},
	//	},
	//	2: {
	//		right: {
	//			face:         3,
	//			newDirection: right,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: c.row,
	//					col: 0,
	//				}
	//			},
	//		},
	//		up: {
	//			face:         1,
	//			newDirection: down,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: 0,
	//					col: faceSize - c.col - 1,
	//				}
	//			},
	//		},
	//		left: {
	//			face:         6,
	//			newDirection: up,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1,
	//					col: faceSize - c.row - 1,
	//				}
	//			},
	//		},
	//		down: {
	//			face:         5,
	//			newDirection: up,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1,
	//					col: faceSize - c.col - 1,
	//				}
	//			},
	//		},
	//	},
	//	3: {
	//		up: {
	//			face:         1,
	//			newDirection: right,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: c.col,
	//					col: 0,
	//				}
	//			},
	//		},
	//		left: {
	//			face:         2,
	//			newDirection: left,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: c.row,
	//					col: faceSize - 1,
	//				}
	//			},
	//		},
	//		right: {
	//			face:         4,
	//			newDirection: right,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: c.row,
	//					col: 0,
	//				}
	//			},
	//		},
	//		down: {
	//			face:         5,
	//			newDirection: right,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1 - c.col,
	//					col: 0,
	//				}
	//			},
	//		},
	//	},
	//	4: {
	//		left: {
	//			face:         3,
	//			newDirection: left,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: c.row,
	//					col: faceSize - 1,
	//				}
	//			},
	//		},
	//		up: {
	//			face:         1,
	//			newDirection: up,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1,
	//					col: c.col,
	//				}
	//			},
	//		},
	//		down: {
	//			face:         5,
	//			newDirection: down,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: 0,
	//					col: c.col,
	//				}
	//			},
	//		},
	//		right: {
	//			face:         6,
	//			newDirection: down,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: 0,
	//					col: faceSize - 1 - c.row,
	//				}
	//			},
	//		},
	//	},
	//	5: {
	//		up: {
	//			face:         4,
	//			newDirection: up,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1,
	//					col: c.col,
	//				}
	//			},
	//		},
	//		right: {
	//			face:         6,
	//			newDirection: right,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: c.row,
	//					col: 0,
	//				}
	//			},
	//		},
	//		down: {
	//			face:         2,
	//			newDirection: up,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1,
	//					col: faceSize - 1 - c.col,
	//				}
	//			},
	//		},
	//		left: {
	//			face:         3,
	//			newDirection: up,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1,
	//					col: faceSize - 1 - c.row,
	//				}
	//			},
	//		},
	//	},
	//	6: {
	//		left: {
	//			face:         5,
	//			newDirection: left,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: c.row,
	//					col: faceSize - 1,
	//				}
	//			},
	//		},
	//		up: {
	//			face:         4,
	//			newDirection: left,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1 - c.col,
	//					col: faceSize - 1,
	//				}
	//			},
	//		},
	//		right: {
	//			face:         1,
	//			newDirection: left,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1 - c.row,
	//					col: faceSize - 1,
	//				}
	//			},
	//		},
	//		down: {
	//			face:         2,
	//			newDirection: right,
	//			transform: func(c coord) coord {
	//				return coord{
	//					row: faceSize - 1 - c.col,
	//					col: 0,
	//				}
	//			},
	//		},
	//	},
	//}

	facesByID := map[int]coord{
		1: {0, 1},
		2: {0, 2},
		3: {1, 1},
		4: {2, 1},
		5: {2, 0},
		6: {3, 0},
	}

	cbe.adjacentFaces = map[int]map[direction]adjacentFace{
		1: {
			up: {
				face:         6,
				newDirection: right,
				transform: func(c coord) coord {
					return coord{
						row: c.col,
						col: 0,
					}
				},
			},
			down: {
				face:         3,
				newDirection: down,
				transform: func(c coord) coord {
					return coord{
						row: 0,
						col: c.col,
					}
				},
			},
			left: {
				face:         5,
				newDirection: right,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1 - c.row,
						col: 0,
					}
				},
			},
			right: {
				face:         2,
				newDirection: right,
				transform: func(c coord) coord {
					return coord{
						row: c.row,
						col: 0,
					}
				},
			},
		},
		2: {
			up: {
				face:         6,
				newDirection: up,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1,
						col: c.col,
					}
				},
			},
			down: {
				face:         3,
				newDirection: left,
				transform: func(c coord) coord {
					return coord{
						row: c.col,
						col: faceSize - 1,
					}
				},
			},
			left: {
				face:         1,
				newDirection: left,
				transform: func(c coord) coord {
					return coord{
						row: c.row,
						col: faceSize - 1,
					}
				},
			},
			right: {
				face:         4,
				newDirection: left,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1 - c.row,
						col: faceSize - 1,
					}
				},
			},
		},
		3: {
			up: {
				face:         1,
				newDirection: up,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1,
						col: c.col,
					}
				},
			},
			down: {
				face:         4,
				newDirection: down,
				transform: func(c coord) coord {
					return coord{
						row: 0,
						col: c.col,
					}
				},
			},
			left: {
				face:         5,
				newDirection: down,
				transform: func(c coord) coord {
					return coord{
						row: 0,
						col: c.row,
					}
				},
			},
			right: {
				face:         2,
				newDirection: up,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1,
						col: c.row,
					}
				},
			},
		},
		4: {
			up: {
				face:         3,
				newDirection: up,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1,
						col: c.col,
					}
				},
			},
			down: {
				face:         6,
				newDirection: left,
				transform: func(c coord) coord {
					return coord{
						row: c.col,
						col: faceSize - 1,
					}
				},
			},
			left: {
				face:         5,
				newDirection: left,
				transform: func(c coord) coord {
					return coord{
						row: c.row,
						col: faceSize - 1,
					}
				},
			},
			right: {
				face:         2,
				newDirection: left,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - c.row - 1,
						col: faceSize - 1,
					}
				},
			},
		},
		5: {
			up: {
				face:         3,
				newDirection: right,
				transform: func(c coord) coord {
					return coord{
						row: c.col,
						col: 0,
					}
				},
			},
			down: {
				face:         6,
				newDirection: down,
				transform: func(c coord) coord {
					return coord{
						row: 0,
						col: c.col,
					}
				},
			},
			left: {
				face:         1,
				newDirection: right,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1 - c.row,
						col: 0,
					}
				},
			},
			right: {
				face:         4,
				newDirection: right,
				transform: func(c coord) coord {
					return coord{
						row: c.row,
						col: 0,
					}
				},
			},
		},
		6: {
			up: {
				face:         5,
				newDirection: up,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1,
						col: c.col,
					}
				},
			},
			down: {
				face:         2,
				newDirection: down,
				transform: func(c coord) coord {
					return coord{
						row: 0,
						col: c.col,
					}
				},
			},
			left: {
				face:         1,
				newDirection: down,
				transform: func(c coord) coord {
					return coord{
						row: 0,
						col: c.row,
					}
				},
			},
			right: {
				face:         4,
				newDirection: up,
				transform: func(c coord) coord {
					return coord{
						row: faceSize - 1,
						col: c.row,
					}
				},
			},
		},
	}

	idsByFace := make(map[coord]int)
	for k, v := range facesByID {
		idsByFace[v] = k
	}

	for c, v := range coords {
		fc := coord{
			row: (c.row - 1) / faceSize,
			col: (c.col - 1) / faceSize,
		}

		faceID := idsByFace[fc]
		if faceID <= 0 {
			fmt.Println(fc)
			panic("unknown face for hardcoded orientation of the cube")
		}

		cc := cubeCoord{
			coord: c,
			face:  faceID,
		}

		if _, ok := cbe.faceRanges[faceID]; !ok {
			cbe.faceRanges[faceID] = faceRange{
				rowRange: coordRange{min: math.MaxInt32, max: 0},
				colRange: coordRange{min: math.MaxInt32, max: 0},
			}
		}

		fr := cbe.faceRanges[faceID]
		if cc.row > fr.rowRange.max {
			fr.rowRange.max = cc.row
		}
		if cc.row < fr.rowRange.min {
			fr.rowRange.min = cc.row
		}
		if cc.col > fr.colRange.max {
			fr.colRange.max = cc.col
		}
		if cc.col < fr.colRange.min {
			fr.colRange.min = cc.col
		}
		cbe.faceRanges[faceID] = fr
		cbe.coords[cc] = v
	}

	return cbe
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	parseInstructions := false
	curRotation := none
	curDigits := ""
	for i, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			parseInstructions = true
			continue
		}

		if !parseInstructions {
			row := i + 1
			if _, ok := rowRanges[row]; !ok {
				rowRanges[row] = coordRange{min: math.MaxInt, max: 0}
			}
			parts := strings.Split(line, "")
			for j, part := range parts {
				col := j + 1
				if _, ok := colRanges[col]; !ok {
					colRanges[col] = coordRange{min: math.MaxInt, max: 0}
				}

				if part == " " {
					continue
				} else if part == "#" || part == "." {
					c := coord{row, col}
					coords[c] = part

					rowRange := rowRanges[row]
					colRange := colRanges[col]
					if col > rowRange.max {
						rowRange.max = col
						rowRanges[row] = rowRange
					}
					if col < rowRange.min {
						rowRange.min = col
						rowRanges[row] = rowRange
					}
					if row > colRange.max {
						colRange.max = row
						colRanges[col] = colRange
					}
					if row < colRange.min {
						colRange.min = row
						colRanges[col] = colRange
					}
				} else {
					panic(part)
				}
			}
		}

		if parseInstructions {
			parts := strings.Split(line, "")
			for _, part := range parts {
				if part == string(clockwise) || part == string(counterClockwise) {
					ins := instruction{rotation: curRotation}
					ins.steps, err = strconv.Atoi(curDigits)
					if err != nil {
						panic(err)
					}
					instructions = append(instructions, ins)
					curRotation = rotation(part)
					curDigits = ""
				} else {
					curDigits += part
				}
			}

			ins := instruction{rotation: curRotation}
			ins.steps, err = strconv.Atoi(curDigits)
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, ins)
		}
	}
}

func getNextDirection(d direction, r rotation) direction {
	if r == none {
		return d
	}

	switch d {
	case right:
		if r == clockwise {
			return down
		} else if r == counterClockwise {
			return up
		} else {
			panic(r)
		}
	case down:
		if r == clockwise {
			return left
		} else if r == counterClockwise {
			return right
		} else {
			panic(r)
		}
	case left:
		if r == clockwise {
			return up
		} else if r == counterClockwise {
			return down
		} else {
			panic(r)
		}
	case up:
		if r == clockwise {
			return right
		} else if r == counterClockwise {
			return left
		} else {
			panic(r)
		}
	default:
		panic(d)
	}
}

func (d direction) asValue() int {
	switch d {
	case right:
		return 0
	case down:
		return 1
	case left:
		return 2
	case up:
		return 3
	default:
		return -1
	}
}
