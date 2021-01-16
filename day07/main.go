package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type bagSpec struct {
	name     string
	contains map[string]int
}

func readInput(filepath string) []bagSpec {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(fileBuffer), "\n")
	r1, _ := regexp.Compile(`(.+) bags contain`)
	r2, _ := regexp.Compile(`(\d+) (.+?) bags?`)

	var rules []bagSpec

	for i, line := range lines {
		if !r1.MatchString(line) {
			fmt.Printf("Invalid format for rule %v (\"%v\")\n", i, line)
			continue
		}

		bagName := r1.FindStringSubmatch(line)[1]

		if !r2.MatchString(line) {
			bag := bagSpec{bagName, nil}
			rules = append(rules, bag)
		} else {
			specs := r2.FindAllStringSubmatch(line, -1)
			m := map[string]int{}
			for _, spec := range specs {
				qty, _ := strconv.Atoi(spec[1])
				m[spec[2]] = qty
			}

			bag := bagSpec{bagName, m}
			rules = append(rules, bag)

		}
	}

	fmt.Printf("Found %v rules\n", len(rules))
	return rules
}

func contains(name string, rules []bagSpec, bags map[string]int) {
	for _, spec := range rules {
		if spec.contains[name] > 0 {
			bags[spec.name]++
			contains(spec.name, rules, bags)
		}
	}
}

func weight(name string, rules []bagSpec) int {
	w := 0
	for _, spec := range rules {
		if spec.name == name {
			if spec.contains == nil {
				return 1
			}

			w = 1
			for bag, qty := range spec.contains {
				w += qty * weight(bag, rules)
			}
		}
	}

	return w
}

func part1(rules []bagSpec) {
	bags := map[string]int{}
	contains("shiny gold", rules, bags)
	fmt.Printf("Part1. %v bags can contains a shiny gold bag\n", len(bags))
}

func part2(rules []bagSpec) {
	weight := weight("shiny gold", rules) - 1 // remove the shiny gold
	fmt.Printf("Part2. %v individual bags inside a shiny gold bag\n", weight)
}

func main() {
	fmt.Println("-- DAY 7 --")

	// rules := readInput("day07/example.txt")
	// rules := readInput("day07/example_part2.txt")
	rules := readInput("day07/input.txt")
	part1(rules)
	part2(rules)
}
