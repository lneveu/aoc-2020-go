package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func readInput(filepath string) []int {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data := bufio.NewReader(strings.NewReader(string(fileBuffer)))
	numbers := []int{}

	for {
		if line, _, err := data.ReadLine(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			n, _ := strconv.Atoi(string(line))
			numbers = append(numbers, n)
		}
	}

	return numbers
}

func part1(adapters []int) {
	sort.Ints(adapters)
	adapters = append([]int{0}, adapters...)                 // prepend 0
	adapters = append(adapters, adapters[len(adapters)-1]+3) // append +3
	diff := map[int]int{1: 0, 2: 0, 3: 0}

	for i := 0; i < len(adapters)-1; i++ {
		diff[adapters[i+1]-adapters[i]]++
	}

	// fmt.Printf("Adapters: %v\n", adapters)
	// fmt.Printf("Diff: %v\n", diff)
	fmt.Printf("Part1. Results: %v\n", diff[1]*diff[3])
}

func part2(adapters []int) {
	sort.Ints(adapters)
	adapters = append([]int{0}, adapters...)                 // prepend 0
	adapters = append(adapters, adapters[len(adapters)-1]+3) // append +3

	a := make([]int, len(adapters))
	a[len(adapters)-1] = 1

	for i := len(adapters) - 2; i >= 0; i-- {
		a[i] = a[i+1]

		if i+3 < len(adapters) && adapters[i+3] <= adapters[i]+3 {
			a[i] += a[i+3]
		}

		if i+2 < len(adapters) && adapters[i+2] <= adapters[i]+3 {
			a[i] += a[i+2]
		}
	}

	// fmt.Printf("Adapters: %v\n", adapters)
	fmt.Printf("Part2. Results: %v\n", a[0])
}

func main() {
	fmt.Println("-- DAY 10 --")

	// inputs := readInput("day10/example.txt")
	inputs := readInput("day10/input.txt")
	part1(inputs)
	part2(inputs)
}
