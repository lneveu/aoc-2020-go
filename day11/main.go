package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type layout struct {
	row, col int
	seats    [][]int
	occupied int
}

const Floor = 0
const Empty = 1
const Occupied = 2

func drawLayout(l layout) {
	for i := 0; i < l.row; i++ {
		for j := 0; j < l.col; j++ {
			switch l.seats[i][j] {
			case Floor:
				fmt.Printf(".") // floor
			case Empty:
				fmt.Printf("L") // empty seat
			case Occupied:
				fmt.Printf("#") // occupied seat
			}
		}
		fmt.Printf("\n")
	}
}

func readInput(filepath string) layout {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data := bufio.NewReader(strings.NewReader(string(fileBuffer)))

	m := layout{row: 0, col: 0, seats: [][]int{{}}, occupied: 0}

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
				m.seats[col] = append(m.seats[col], Floor)
			} else if char == "L" {
				m.seats[col] = append(m.seats[col], Empty)
			} else if char == "#" {
				m.seats[col] = append(m.seats[col], Occupied)
			} else if char == "\n" {
				m.seats = append(m.seats, []int{})
				col++
			}
		}
	}

	m.row = len(m.seats)
	m.col = len(m.seats[0])

	return m
}

func checkIfNoOccupiedAdjacent(l layout, seatRow int, seatCol int) bool {
	for r := seatRow - 1; r <= seatRow+1; r++ {
		for c := seatCol - 1; c <= seatCol+1; c++ {
			if r >= l.row || c >= l.col || r < 0 || c < 0 || (r == seatRow && c == seatCol) {
				continue
			}
			if l.seats[r][c] == Occupied {
				return false
			}
		}
	}
	return true
}

func checkIfFourOrMoreOccupiedAdjacent(l layout, seatRow int, seatCol int) bool {
	occupied := 0
	for r := seatRow - 1; r <= seatRow+1; r++ {
		for c := seatCol - 1; c <= seatCol+1; c++ {
			if r >= l.row || c >= l.col || r < 0 || c < 0 || (r == seatRow && c == seatCol) {
				continue
			}
			if l.seats[r][c] == Occupied {
				occupied++
			}

			if occupied >= 4 {
				return true
			}
		}
	}
	return false
}

func checkIfNoOccupiedVisible(l layout, seatRow int, seatCol int) bool {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}

			row := seatRow
			col := seatCol

			for {
				row += x
				col += y

				if row >= l.row || col >= l.col || row < 0 || col < 0 || l.seats[row][col] == Empty {
					break
				}

				if l.seats[row][col] == Occupied {
					return false
				}
			}
		}
	}
	return true
}

func checkIfFiveOrMoreOccupiedVisible(l layout, seatRow int, seatCol int) bool {
	occupied := 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}

			row := seatRow
			col := seatCol

			for {
				row += x
				col += y
				if row >= l.row || col >= l.col || row < 0 || col < 0 || l.seats[row][col] == Empty {
					break
				}

				if l.seats[row][col] == Occupied {
					occupied++
					break
				}
			}

			if occupied >= 5 {
				return true
			}
		}
	}
	return false
}

func copySeats(seats [][]int) [][]int {
	var cop = make([][]int, len(seats))
	copy(cop, seats)
	for i := range cop {
		cop[i] = make([]int, len(seats[i]))
		copy(cop[i], seats[i])
	}
	return cop
}

func applyRulesPart1(l *layout) {
	seats := copySeats(l.seats)

	for i := 0; i < l.row; i++ {
		for j := 0; j < l.col; j++ {
			seat := l.seats[i][j]

			if seat == Empty && checkIfNoOccupiedAdjacent(*l, i, j) {
				seats[i][j] = Occupied
				l.occupied++
			}

			if seat == Occupied && checkIfFourOrMoreOccupiedAdjacent(*l, i, j) {
				seats[i][j] = Empty
				l.occupied--
			}
		}
	}
	l.seats = seats
}

func applyRulesPart2(l *layout) {
	seats := copySeats(l.seats)

	for i := 0; i < l.row; i++ {
		for j := 0; j < l.col; j++ {
			seat := l.seats[i][j]

			if seat == Empty && checkIfNoOccupiedVisible(*l, i, j) {
				seats[i][j] = Occupied
				l.occupied++
			}

			if seat == Occupied && checkIfFiveOrMoreOccupiedVisible(*l, i, j) {
				seats[i][j] = Empty
				l.occupied--
			}
		}
	}
	l.seats = seats
}

func part1(l layout) {
	prevOccupiedSeats := -1
	for {
		applyRulesPart1(&l)
		if l.occupied == prevOccupiedSeats {
			fmt.Printf("Part1. %v occupied seats\n", l.occupied)
			return
		}
		prevOccupiedSeats = l.occupied
	}
}

func part2(l layout) {
	prevOccupiedSeats := -1
	for {
		applyRulesPart2(&l)
		if l.occupied == prevOccupiedSeats {
			fmt.Printf("Part2. %v occupied seats\n", l.occupied)
			return
		}
		prevOccupiedSeats = l.occupied
	}
}

func main() {
	fmt.Println("-- DAY 11 --")

	// layout := readInput("day11/example.txt")
	layout := readInput("day11/input.txt")
	// drawLayout(layout)
	part1(layout)
	part2(layout)
}
