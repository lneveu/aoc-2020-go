package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	v1, v2   int
	letter   rune
	password string
}

func readInput(filepath string) []passwordPolicy {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	list := strings.Split(string(content), "\n")
	r, _ := regexp.Compile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)

	var policies []passwordPolicy

	for i, line := range list {
		if !r.MatchString(line) {
			fmt.Printf("Invalid format for line %v (\"%v\")\n", i, line)
			continue
		}
		matches := r.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			v1, _ := strconv.Atoi(m[1])
			v2, _ := strconv.Atoi(m[2])
			letter := rune(m[3][0])

			policy := passwordPolicy{v1: v1, v2: v2, letter: letter, password: m[4]}
			policies = append(policies, policy)
		}
	}

	return policies
}

func part1(inputs []passwordPolicy) {
	validPassword := 0

	for _, policy := range inputs {
		sum := 0

		for _, c := range policy.password {
			if c == policy.letter {
				sum++
			}
		}

		min := policy.v1
		max := policy.v2
		if sum >= min && sum <= max {
			validPassword++
		}
	}

	fmt.Printf("Part1. Number of valid passwords: %v\n", validPassword)
}

func part2(inputs []passwordPolicy) {
	validPassword := 0

	for _, policy := range inputs {
		pos1Letter := rune(policy.password[policy.v1-1])
		pos2Letter := rune(policy.password[policy.v2-1])
		policyLetter := policy.letter

		if (pos1Letter == policyLetter) != (pos2Letter == policyLetter) {
			validPassword++
		}

	}

	fmt.Printf("Part2. Number of valid passwords: %v\n", validPassword)
}

func main() {
	fmt.Println("-- DAY 2 --")

	// inputs := readInput("day02/example.txt")
	inputs := readInput("day02/input.txt")
	part1(inputs)
	part2(inputs)
}
