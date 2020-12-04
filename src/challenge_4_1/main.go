package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const FileName = "input_4"

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
	valid bool
}

func main() {
	dir, err := filepath.Abs("src/inputs/" + FileName)
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	var passportText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			lines = append(lines, passportText)
			passportText = ""
		} else {
			//add all text to passportText until an empty line is hit
			passportText += scanner.Text() + " "
		}
	}
	lines = append(lines, passportText)

	var passports []Passport
	for _, line := range lines {
		passports = append(passports, makePassport(line))
	}
	totalValidPassports := 0

	for _, pass := range passports {
		if pass.valid == true {
			totalValidPassports++
		}
	}

	fmt.Println(totalValidPassports)
}

func makePassport(passportText string) Passport {
	fields := strings.Split(passportText, " ")
	fieldMap := make(map[string]string)
	for _, field := range fields {
		data := strings.Split(field, ":")
		if data[0] == "" {
			continue
		}
		fieldMap[data[0]] = data[1]
	}

	return Passport{
		byr:   fieldMap["byr"],
		iyr:   fieldMap["iyr"],
		eyr:   fieldMap["eyr"],
		hgt:   fieldMap["hgt"],
		hcl:   fieldMap["hcl"],
		ecl:   fieldMap["ecl"],
		pid:   fieldMap["pid"],
		cid:   fieldMap["cid"],
		valid: valid(fieldMap),
	}
}

func valid(fieldMap map[string]string) bool {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid",
	}

	eyeColors := map[string]string{
		"amb": "amb",
		"blu": "blu",
		"brn": "brn",
		"gry": "gry",
		"grn": "grn",
		"hzl": "hzl",
		"oth": "oth",
	}

	//correct fields
	for _, reqField := range requiredFields {
		hasField := fieldMap[reqField]
		if hasField == "" {
			return false
		}
	}

	//validate byr
	byr, err := strconv.Atoi(fieldMap["byr"])
	if err != nil {
		return false
	}
	if !between(byr, 1920, 2002) {
		return false
	}
	//validate iyr
	iyr, err := strconv.Atoi(fieldMap["iyr"])

	if err != nil {
		return false
	}
	if !between(iyr, 2010, 2020) {
		return false
	}

	//validate eyr
	eyr, err := strconv.Atoi(fieldMap["eyr"])
	if err != nil {
		return false
	}
	if !between(eyr, 2020, 2030) {
		return false
	}

	//validate hgt
	hgt := fieldMap["hgt"]
	isCm := strings.Contains(hgt, "cm")
	isIn := strings.Contains(hgt, "in")
	if !isCm && !isIn {
		return false
	}
	convertedHgt, _ := strconv.Atoi(hgt[0:len(hgt) - 2])
	if isCm {

		if !between(convertedHgt, 150, 193) {
			return false
		}
	}
	if isIn {
		if !between(convertedHgt, 59, 76) {
			return false
		}
	}

	//validate hcl
	rHcl := regexp.MustCompile("#[a-fA-F0-9]{6}")
	if !rHcl.MatchString(fieldMap["hcl"]) {
		return false
	}

	//validate ecl
	eyeColor := eyeColors[fieldMap["ecl"]]
	if eyeColor == "" {
		return false
	}

	//validate hcl
	rPid := regexp.MustCompile("\\d{9}")
	if !rPid.MatchString(fieldMap["pid"]) {
		return false
	}

	return true
}

func between(number int, min int, max int) bool {
	return number >= min && number <= max
}
