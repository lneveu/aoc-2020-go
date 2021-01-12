package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func readInput(filepath string) []string {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data := bufio.NewReader(strings.NewReader(string(fileBuffer)))

	inputs := []string{}

	for {
		if line, _, err := data.ReadLine(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			inputs = append(inputs, string(line))
		}
	}
	fmt.Printf("Found %v boarding passes\n", len(inputs))
	return inputs
}

func getSeatID(boardPass string) int {
	minRow := 0
	maxRow := 127
	minSeat := 0
	maxSeat := 7

	// row
	for _, c := range boardPass[:7] {
		if string(c) == "F" {
			maxRow = maxRow - (maxRow-minRow+1)/2
		}

		if string(c) == "B" {
			minRow = minRow + (maxRow-minRow+1)/2
		}
	}

	// seat
	for _, c := range boardPass[7:] {
		if string(c) == "L" {
			maxSeat = maxSeat - (maxSeat-minSeat+1)/2
		}

		if string(c) == "R" {
			minSeat = minSeat + (maxSeat-minSeat+1)/2
		}
	}

	seatID := minRow*8 + minSeat
	// fmt.Printf("ROW %v SEAT %v SEATID %d (boardPass = %v)\n", minRow, minSeat, seatID, boardPass)

	return seatID
}

func part1(inputs []string) {
	maxSeatID := 0
	for _, boardPass := range inputs {
		seatID := getSeatID(boardPass)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Printf("Part1. Highest seat ID: %v\n", maxSeatID)
}

func part2(inputs []string) {
	seats := []int{}
	for _, boardPass := range inputs {
		seatID := getSeatID(boardPass)
		seats = append(seats, seatID)
	}

	sort.Ints(seats)

	mySeat := 0

	for i := range seats {
		if i > (len(seats) - 2) {
			break
		}

		if seats[i+1]-seats[i] > 1 {
			mySeat = seats[i] + 1
			break
		}
	}

	fmt.Printf("Part2. Our seat ID is %v \n", mySeat)
}

func main() {
	fmt.Println("-- DAY 5 --")

	// inputs := readInput("day05/example.txt")
	inputs := readInput("day05/input.txt")
	part1(inputs)
	part2(inputs)
}
