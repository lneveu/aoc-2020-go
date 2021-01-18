package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

func part1(numbers []int, preamble int) int {
	i := 0
	for i = preamble; i < len(numbers)-1; i++ {
		found := false
		for j := i - preamble; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if numbers[j]+numbers[k] == numbers[i] {
					found = true
					break
				}
			}
		}

		if !found {
			fmt.Printf("Part1. Invalid number: %v\n", numbers[i])
			break
		}
	}
	return numbers[i]
}

const MaxInt = int(^uint(0) >> 1)

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func part2(target int, numbers []int, preamble int) {
	found := false
	idx := 0

	sum := 0
	min := MaxInt
	max := 0

	for !found {
		if idx > len(numbers) {
			break
		}

		for i := idx; i < len(numbers); i++ {
			sum += numbers[i]
			min = Min(min, numbers[i])
			max = Max(max, numbers[i])

			if sum > target {
				// reset
				idx++
				sum = 0
				min = MaxInt
				max = 0
				break
			}

			if sum == target && i >= 2 {
				found = true
				break
			}
		}
	}

	fmt.Printf("Part2. Encryption weakness: %v\n", min+max)
}

func main() {
	fmt.Println("-- DAY 9 --")

	// inputs := readInput("day09/example.txt")
	// res := part1(inputs, 5)
	// part2(res, inputs, 5)

	inputs := readInput("day09/input.txt")
	res := part1(inputs, 25)
	part2(res, inputs, 25)
}
