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

type operation struct {
	code string
	arg  int
	exec bool
}
type program struct {
	acc int
	ops []operation
}

func readInput(filepath string) program {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data := bufio.NewReader(strings.NewReader(string(fileBuffer)))
	ops := []operation{}

	for {
		if line, _, err := data.ReadLine(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			opStr := strings.Split(string(line), " ")
			arg, _ := strconv.Atoi(opStr[1])
			ops = append(ops, operation{code: opStr[0], arg: arg})
		}
	}

	return program{0, ops}
}

func run(p *program) bool {
	idx := 0
	success := false

	for {
		if idx == len(p.ops) {
			fmt.Printf("Program ended successfully\n")
			success = true
			break
		}

		op := &p.ops[idx]

		if op.exec == true {
			// fmt.Printf("Infinite loop (last instruction: %v acc: %v)\n", p.ops[idx].code, p.acc)
			success = false
			break
		}

		op.exec = true
		switch op.code {
		case "nop":
			idx++
			break
		case "acc":
			idx++
			p.acc += op.arg
			break
		case "jmp":
			idx += op.arg
			break
		}
	}

	return success
}

func copyProgram(p program) program {
	var ops = make([]operation, len(p.ops))
	copy(ops, p.ops)
	return program{p.acc, ops}
}

func part1(p program) {
	run(&p)
	fmt.Printf("Part1. accumulator value: %v\n", p.acc)
}

func part2(p program) {
	permIdx := 0

	for {
		if permIdx > len(p.ops)-1 {
			fmt.Printf("Can't repair the program\n")
			break
		}

		if p.ops[permIdx].code == "acc" {
			permIdx++
			continue
		}

		pCopy := copyProgram(p)

		if p.ops[permIdx].code == "nop" {
			(&pCopy.ops[permIdx]).code = "jmp"
		} else if p.ops[permIdx].code == "jmp" {
			(&pCopy.ops[permIdx]).code = "nop"
		}

		success := run(&pCopy)
		if success {
			fmt.Printf("Part2. Program repaired! Permutation idx: %v - accumulator value: %v\n", permIdx, pCopy.acc)
			break
		}
		permIdx++
	}
}

func main() {
	fmt.Println("-- DAY 8 --")

	// program := readInput("day08/example.txt")
	program := readInput("day08/input.txt")
	// part1(program)
	part2(program)
}
