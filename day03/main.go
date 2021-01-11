package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type treesMap struct {
	row, col int
	val      [][]int
}

func drawMap(m treesMap) {
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			switch m.val[i][j] {
			case 0:
				fmt.Printf(".")
			case 1:
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
}

func readInput(filepath string) treesMap {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data := bufio.NewReader(strings.NewReader(string(fileBuffer)))

	m := treesMap{row: 0, col: 0, val: [][]int{{}}}

	col := 0

	for {
		if c, _, err := data.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			char := string(c)

			if char == "." {
				m.val[col] = append(m.val[col], 0)
			} else if char == "#" {
				m.val[col] = append(m.val[col], 1)
			} else if char == "\n" {
				m.val = append(m.val, []int{})
				col++
			}
		}
	}

	m.row = len(m.val)
	m.col = len(m.val[0])

	return m
}

func traverse(slope []int, m treesMap) int {
	treesCount := 0
	posCol := 0
	posRow := 0

	for {
		posCol += slope[0]
		posRow += slope[1]

		if posCol >= m.col {
			posCol = posCol % m.col
		}

		if posRow >= m.row {
			break
		}

		if m.val[posRow][posCol] == 1 {
			treesCount++
		}
	}

	return treesCount
}

func part1(m treesMap) {
	result := traverse([]int{3, 1}, m)

	fmt.Printf("Part1. Encouter %v trees\n", result)
}

func part2(m treesMap) {
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	hits := []int{}
	result := 1
	for _, slope := range slopes {
		h := traverse(slope, m)
		hits = append(hits, h)
		result *= h
	}

	fmt.Printf("Part2. (found %v trees) - Results=%v\n", hits, result)
}

func main() {
	fmt.Println("-- DAY 3 --")

	// inputs := readInput("day03/example.txt")
	m := readInput("day03/input.txt")
	// drawMap(m)
	part1(m)
	part2(m)
}
