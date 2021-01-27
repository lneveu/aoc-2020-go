package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type ShipPos struct {
	x     int
	y     int
	angle int // 0 = N; 90=E; 180=S; 270=W
}

type WaypointPos struct {
	x int
	y int
}

type Move struct {
	action string
	val    int
}

func readInput(filepath string) []Move {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data := bufio.NewReader(strings.NewReader(string(fileBuffer)))
	r, _ := regexp.Compile(`^(N|S|E|W|L|R|F)(\d+)$`)
	moves := []Move{}

	for {
		if line, _, err := data.ReadLine(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {

			matches := r.FindAllStringSubmatch(string(line), -1)
			for _, m := range matches {
				action := m[1]
				value, _ := strconv.Atoi(m[2])
				moves = append(moves, Move{action, value})
			}
		}
	}

	return moves
}

func part1(moves []Move) {
	shipPos := ShipPos{0, 0, 90}

	for _, move := range moves {
		switch move.action {
		case "N":
			shipPos.y += move.val
			break
		case "S":
			shipPos.y -= move.val
			break
		case "E":
			shipPos.x += move.val
			break
		case "W":
			shipPos.x -= move.val
			break
		case "L":
			shipPos.angle = (shipPos.angle - move.val + 360) % 360
			break
		case "R":
			shipPos.angle = (shipPos.angle + move.val) % 360
			break
		case "F":
			switch shipPos.angle {
			case 0:
				shipPos.y += move.val
				break
			case 90:
				shipPos.x += move.val
				break
			case 180:
				shipPos.y -= move.val
				break
			case 270:
				shipPos.x -= move.val
				break
			}
			break
		}
	}

	fmt.Printf("Part1. ship position: (x:%v y:%v angle:%v)\n", shipPos.x, shipPos.y, shipPos.angle)

	dist := Abs(shipPos.x) + Abs(shipPos.y)
	fmt.Printf("Part1. distance: %v\n", dist)
}

func part2(moves []Move) {
	shipPos := ShipPos{0, 0, 90}
	waypointPos := WaypointPos{10, 1} // relative position

	for _, move := range moves {
		switch move.action {
		case "N":
			waypointPos.y += move.val
			break
		case "S":
			waypointPos.y -= move.val
			break
		case "E":
			waypointPos.x += move.val
			break
		case "W":
			waypointPos.x -= move.val
			break
		case "L":
			switch move.val {
			case 90:
				old := waypointPos.x
				waypointPos.x = -waypointPos.y
				waypointPos.y = old
				break
			case 180:
				waypointPos.x = -waypointPos.x
				waypointPos.y = -waypointPos.y
				break
			case 270:
				old := waypointPos.x
				waypointPos.x = waypointPos.y
				waypointPos.y = -old
				break
			}
			break
		case "R":
			switch move.val {
			case 90:
				old := waypointPos.x
				waypointPos.x = waypointPos.y
				waypointPos.y = -old
				break
			case 180:
				waypointPos.x = -waypointPos.x
				waypointPos.y = -waypointPos.y
				break
			case 270:
				old := waypointPos.x
				waypointPos.x = -waypointPos.y
				waypointPos.y = old
				break
			}
			break
		case "F":
			shipPos.x += waypointPos.x * move.val
			shipPos.y += waypointPos.y * move.val
			break
		}

	}

	fmt.Printf("Part2. ship position: (x:%v y:%v)\n", shipPos.x, shipPos.y)
	fmt.Printf("Part2. waypoint position: (x:%v y:%v)\n", waypointPos.x, waypointPos.y)

	dist := Abs(shipPos.x) + Abs(shipPos.y)
	fmt.Printf("Part2. distance: %v\n", dist)
}

func main() {
	fmt.Println("-- DAY 12 --")

	// moves := readInput("day12/example.txt")
	moves := readInput("day12/input.txt")
	part1(moves)
	part2(moves)
}
