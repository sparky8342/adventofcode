package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

func get_sizes(top *Dir) int {
	sum := 0
	_ = find_sizes(top, &sum)
	return sum
}

func find_sizes(dir *Dir, sum *int) int {
	total := 0
	for _, file := range dir.files {
		total += file.size
	}
	for _, sub_dir := range dir.dirs {
		total += find_sizes(sub_dir, sum)
	}
	if total <= 100000 {
		*sum += total
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
	size := get_sizes(tree)
	fmt.Println(size)
}
