package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	//correct fields
	for _, reqField := range requiredFields {
		hasField := fieldMap[reqField]
		if hasField == "" {
			return false
		}
	}

	return true
}