package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type snafu []int

var (
	numbers []snafu
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	sum := 0
	for _, num := range numbers {
		sum += num.ToDecimal()
	}

	sumSNAFU := ToSNAFU(sum).ToString()
	fmt.Printf("%v\n", sumSNAFU)
}

func part2() {
	fmt.Printf("%v\n", "N/A")
}

func (n snafu) ToDecimal() int {
	total := 0
	for place, digit := range n {
		decimalPlaceValue := 1
		for i := 0; i < place; i++ {
			decimalPlaceValue *= 5
		}

		total += decimalPlaceValue * digit
	}

	return total
}

func (n snafu) ToString() string {
	s := ""
	for place := len(n) - 1; place >= 0; place-- {
		digit := n[place]
		switch digit {
		case 0:
			s += "0"
		case 1:
			s += "1"
		case 2:
			s += "2"
		case -1:
			s += "-"
		case -2:
			s += "="
		default:
			panic(digit)
		}
	}

	return s
}

func ToSNAFU(decimal int) snafu {
	const base = 5
	highestPower := int(math.Log(float64(decimal)) / math.Log(base))
	digits := make([]int, highestPower+1)
	r := decimal
	for highestPower >= 0 {
		power := int(math.Round(math.Pow(base, float64(highestPower))))
		q := r / power
		r = r % power
		digits[highestPower] = q
		highestPower--
	}

	for place, digit := range digits {
		if digit > base || digit < 0 {
			panic("bad digit")
		}

		if digit == 3 {
			if place == len(digits)-1 {
				digits = append(digits, 0)
			}
			digits[place+1]++
			digits[place] = -2
		}

		if digit == 4 {
			if place == len(digits)-1 {
				digits = append(digits, 0)
			}
			digits[place+1]++
			digits[place] = -1
		}

		if digit == 5 {
			digits[place+1]++
			digits[place] = 0
		}
	}

	return digits
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, "")
		var num snafu
		for place := len(parts) - 1; place >= 0; place-- {
			digitStr := parts[place]
			var digit int
			switch digitStr {
			case "0":
				digit = 0
			case "1":
				digit = 1
			case "2":
				digit = 2
			case "-":
				digit = -1
			case "=":
				digit = -2
			}

			num = append(num, digit)
		}

		numbers = append(numbers, num)
	}
}
