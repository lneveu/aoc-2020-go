package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(filepath string) []int {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	arr := strings.Split(string(content), "\n")

	var inputs = []int{}
	for _, i := range arr {
		n, _ := strconv.Atoi(i)
		inputs = append(inputs, n)
	}

	return inputs
}

func findPairThatSum2020(arr []int) []int {
	for _, val1 := range arr {
		for _, val2 := range arr {
			if val1 == val2 {
				continue
			}
			if val1+val2 == 2020 {
				return []int{val1, val2}
			}
		}
	}
	return nil
}

// Find every unique combinations for a given array
// Size of the combination tuples can be specified
func combinations(arr []int, size int) [][]int {
	var comb = [][]int{}

	var f func(int, []int, []int)
	f = func(s int, src []int, acc []int) {
		if s == 0 {
			if len(acc) > 0 {
				comb = append(comb, acc)
			}
			return
		}

		for i := 0; i < len(src); i++ {
			f(s-1, src[i+1:], append(acc, src[i]))
		}
	}

	f(size, arr, []int{})
	return comb
}

func part1(inputs []int) {
	results := findPairThatSum2020(inputs)

	if results != nil {
		fmt.Printf("Part1. Found a matching pair: %v - Result=%v\n", results, results[0]*results[1])
	} else {
		fmt.Println("Part1. No pair found.")
	}
}

func part2(inputs []int) {
	combinations := combinations(inputs, 3)

	var results []int = nil

	// Iterate over every combinations and find that one sums 2020
	for _, c := range combinations {
		if c[0]+c[1]+c[2] == 2020 {
			results = c
			break
		}
	}

	if results != nil {
		fmt.Printf("Part2. Found a matching tuple: %v - Result=%v\n", results, results[0]*results[1]*results[2])
	} else {
		fmt.Println("Part2. No tuple found.")
	}
}

func main() {
	fmt.Println("-- DAY 1 --")

	inputs := readInput("day01/input.txt")
	part1(inputs)
	part2(inputs)
}
