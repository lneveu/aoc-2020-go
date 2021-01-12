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

// passport type
type passport struct {
	byr int    // Birth Year
	iyr int    // Issue Year
	eyr int    // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

func (p passport) isValid() (bool, string) {

	if p.byr == 0 {
		return false, "Missing birth year (byr)"
	}

	if p.iyr == 0 {
		return false, "Missing issue year (iyr)"
	}

	if p.eyr == 0 {
		return false, "Missing expiration year (eyr)"
	}

	if p.hgt == "" {
		return false, "Missing height (hgt)"
	}

	if p.hcl == "" {
		return false, "Missing hair color (hcl)"
	}

	if p.ecl == "" {
		return false, "Missing eye color (ecl)"
	}

	if p.pid == "" {
		return false, "Missing passport id (pid)"
	}

	// Coutry ID is optional

	return true, ""
}

var heightReg = regexp.MustCompile(`^(\d+)(cm|in)$`)
var hairColorReg = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var eyeColorReg = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
var pidReg = regexp.MustCompile(`^\d{9}$`)

func validateHeight(height string) (bool, string) {

	if !heightReg.MatchString(height) {
		return false, "Wrong format"
	}

	matches := heightReg.FindAllStringSubmatch(height, -1)
	for _, m := range matches {
		value, err := strconv.Atoi(m[1])
		metric := m[2]

		if err != nil {
			return false, "Invalid height"
		}

		switch metric {
		case "cm":
			if value < 150 || value > 193 {
				return false, fmt.Sprintf("Height out of band - must be > 150cm and < 193cm (%v)", height)
			}
		case "in":
			if value < 59 || value > 76 {
				return false, fmt.Sprintf("Height out of band - must be > 59in and < 76in (%v)", height)
			}
		default:
			return false, "Invalid metric (must be cm or in)"
		}

	}

	return true, ""
}

func (p passport) isValidStrict() (bool, string) {

	if p.byr == 0 {
		return false, "Missing birth year (byr)"
	}

	if p.byr < 1920 || p.byr > 2002 {
		return false, fmt.Sprintf("Birth year out of bounds (%v)", p.byr)
	}

	if p.iyr == 0 {
		return false, "Missing issue year (iyr)"
	}

	if p.iyr < 2010 || p.iyr > 2020 {
		return false, fmt.Sprintf("Issue year out of bounds (%v)", p.iyr)
	}

	if p.eyr == 0 {
		return false, "Missing expiration year (eyr)"
	}

	if p.eyr < 2020 || p.eyr > 2030 {
		return false, fmt.Sprintf("Expiration year out of bounds (%v)", p.eyr)
	}

	if p.hgt == "" {
		return false, "Missing height (hgt)"
	}

	valid, err := validateHeight(p.hgt)

	if !valid {
		return false, fmt.Sprintf("Invalid height (%v)", err)
	}

	if p.hcl == "" {
		return false, "Missing hair color (hcl)"
	}

	if !hairColorReg.Match([]byte(p.hcl)) {
		return false, fmt.Sprintf("Invalid hair color (%v)", p.hcl)
	}

	if p.ecl == "" {
		return false, "Missing eye color (ecl)"
	}

	if !eyeColorReg.Match([]byte(p.ecl)) {
		return false, fmt.Sprintf("Invalid eye color (%v)", p.ecl)
	}

	if p.pid == "" {
		return false, "Missing passport id (pid)"
	}

	if !pidReg.Match([]byte(p.pid)) {
		return false, fmt.Sprintf("Invalid PID (%v)", p.pid)
	}

	// Coutry ID is optional

	return true, ""
}

func readInput(filepath string) []passport {
	fileBuffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data := bufio.NewReader(strings.NewReader(string(fileBuffer)))

	passports := []passport{}
	curPassport := passport{}

	for {
		if line, _, err := data.ReadLine(); err != nil {
			if err == io.EOF {
				// Dont forget the last one
				passports = append(passports, curPassport)
				break
			} else {
				log.Fatal(err)
			}
		} else {
			strLine := string(line)
			if strLine == "" {
				passports = append(passports, curPassport)
				curPassport = passport{}
			} else {
				fields := strings.Fields(strLine)
				for _, field := range fields {
					kv := strings.Split(field, ":")
					key := kv[0]
					value := kv[1]

					switch key {
					case "byr":
						i, _ := strconv.Atoi(value)
						curPassport.byr = i
					case "iyr":
						i, _ := strconv.Atoi(value)
						curPassport.iyr = i
					case "eyr":
						i, _ := strconv.Atoi(value)
						curPassport.eyr = i
					case "hgt":
						curPassport.hgt = value
					case "hcl":
						curPassport.hcl = value
					case "ecl":
						curPassport.ecl = value
					case "pid":
						curPassport.pid = value
					case "cid":
						curPassport.cid = value
					}
				}
			}
		}
	}
	fmt.Printf("Found %v passports\n", len(passports))
	return passports
}

func part1(passports []passport) {
	validPassports := 0
	for _, p := range passports {
		valid, err := p.isValid()
		if valid {
			validPassports++
		} else {
			fmt.Printf("Invalid passport: %v\n", err)
		}
	}

	fmt.Printf("Part1. %v valid passports\n", validPassports)
}

func part2(passports []passport) {
	validPassports := 0
	for _, p := range passports {
		valid, err := p.isValidStrict()
		if valid {
			validPassports++
		} else {
			fmt.Printf("Invalid passport: %v\n", err)
		}
	}

	fmt.Printf("Part2. %v valid passports\n", validPassports)
}

// TODO: use a dedicated grammar and parser
func main() {
	fmt.Println("-- DAY 4 --")

	// passports := readInput("day04/example.txt")
	// passports := readInput("day04/example_invalid.txt")
	// passports := readInput("day04/example_valid.txt")
	passports := readInput("day04/input.txt")
	part1(passports)
	part2(passports)
}
