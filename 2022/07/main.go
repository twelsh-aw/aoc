package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	subdirectories []*directory
	files          []*file
}

var filesystem = make(map[string]*directory)
var cwd string

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, v := range filesystem {
		size := getSize(v)
		if size <= 100000 {
			total += size
		}
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	const maxSize = 70000000
	const updateSize = 30000000
	currentSize := getSize(filesystem["/"])
	spaceRequired := currentSize + updateSize - maxSize
	toDelete := currentSize
	if spaceRequired <= 0 {
		toDelete = 0
	} else {
		for _, v := range filesystem {
			size := getSize(v)
			if size >= spaceRequired && size < toDelete {
				toDelete = size
			}
		}
	}
	fmt.Printf("%v\n", toDelete)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					cwd = "/"
				} else if parts[2] == ".." {
					cwdParts := strings.Split(cwd, "/")
					cwd = strings.Join(cwdParts[0:len(cwdParts)-2], "/") + "/"
				} else {
					cwd = cwd + parts[2] + "/"
				}
			}

			if parts[1] == "ls" {
				// is ls twice in a row possible?
				continue
			}

			continue
		}

		if filesystem[cwd] == nil {
			filesystem[cwd] = &directory{}
		}
		dir := filesystem[cwd]

		if parts[0] == "dir" {
			subdir := cwd + parts[1] + "/"
			if filesystem[subdir] == nil {
				filesystem[subdir] = &directory{}
			}
			dir.subdirectories = append(dir.subdirectories, filesystem[subdir])
			continue
		}

		size, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Sprintf("error parsing file size: %s", err))
		}

		f := file{
			name: parts[1],
			size: size,
		}
		dir.files = append(dir.files, &f)
	}
}

func getSize(d *directory) int {
	size := 0
	for _, f := range d.files {
		size += f.size
	}
	for _, s := range d.subdirectories {
		size += getSize(s)
	}

	return size
}

func printFileSystem() {
	for k, v := range filesystem {
		for _, f := range v.files {
			fmt.Printf("FS %s: %+v\n", k, f)
		}
	}
}
