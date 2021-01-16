package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func readInput(filepath string) [][]string {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data := bufio.NewReader(strings.NewReader(string(fileBuffer)))

	groupIdx := 0
	inputs := [][]string{{}}

	for {
		if line, _, err := data.ReadLine(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if string(line) == "" {
				inputs = append(inputs, []string{})
				groupIdx++
			} else {
				inputs[groupIdx] = append(inputs[groupIdx], string(line))
			}
		}
	}

	fmt.Printf("Found %v groups\n", groupIdx+1)
	return inputs
}

func part1(groups [][]string) {
	sum := 0
	for _, g := range groups {
		m := map[rune]int{}
		for _, p := range g {
			for _, ans := range p {
				m[ans]++
			}
		}
		sum += len(m)
	}
	fmt.Printf("Part1. Sum of yes: %v\n", sum)
}

func part2(groups [][]string) {
	total := 0
	for _, g := range groups {
		m := map[rune]int{}
		for _, p := range g {
			for _, ans := range p {
				m[ans]++
			}
		}

		nbPer := len(g)

		sum := 0
		for _, v := range m {
			if v == nbPer {
				sum++
			}
		}
		total += sum
	}
	fmt.Printf("Part2. Total of questions which everyone said yes %v \n", total)
}

func main() {
	fmt.Println("-- DAY 6 --")

	// groups := readInput("day06/example.txt")
	groups := readInput("day06/input.txt")
	part1(groups)
	part2(groups)
}
