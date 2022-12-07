package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const MAX_DIR_SIZE = 100000
const TOTAL_DISK = 70000000
const SPACE_NEEDED = 30000000

type File struct {
	name string
	size int
}

type Dir struct {
	dirs   map[string]*Dir
	parent *Dir
	files  []File
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func get_sizes(top *Dir) (int, int) {
	sum := 0
	smallest := 0
	total := find_sizes(top, &sum, 0, &smallest)

	sum = 0
	free := TOTAL_DISK - total
	needed := SPACE_NEEDED - free
	smallest = math.MaxInt32
	_ = find_sizes(top, &sum, needed, &smallest)
	return sum, smallest
}

func find_sizes(dir *Dir, sum *int, needed int, smallest *int) int {
	total := 0
	for _, file := range dir.files {
		total += file.size
	}
	for _, sub_dir := range dir.dirs {
		total += find_sizes(sub_dir, sum, needed, smallest)
	}
	if total <= MAX_DIR_SIZE {
		*sum += total
	}
	if total >= needed && total < *smallest {
		*smallest = total
	}
	return total
}

func create_tree(data []string) *Dir {
	top := &Dir{dirs: map[string]*Dir{}}
	node := top

	for _, line := range data {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					node = top
				} else if parts[2] == ".." {
					node = node.parent
				} else {
					node = node.dirs[parts[2]]
				}
			}
		} else if parts[0] == "dir" {
			node.dirs[parts[1]] = &Dir{parent: node, dirs: map[string]*Dir{}}
		} else {
			size, _ := strconv.Atoi(parts[0])
			node.files = append(node.files, File{name: parts[1], size: size})
		}
	}
	return top
}

func main() {
	data := load_data("input.txt")

	tree := create_tree(data)
	size, smallest := get_sizes(tree)
	fmt.Println(size)
	fmt.Println(smallest)
}
