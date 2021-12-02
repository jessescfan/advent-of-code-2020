package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const FileName = "input_7"

func main() {
	dir, err := filepath.Abs("src/inputs/" + FileName)
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	bagCollection := map[string][]string{}
	shinyGoldHolders := map[string]string{}
	for _, line := range lines {
		bag := strings.Split(line, "contain")
		bagNameT := strings.Split(bag[0], "bag")
		bagName := strings.TrimSpace(bagNameT[0])
		bagRules := bag[1]
		var ruleBagNames []string
		rules := strings.Split(bagRules, ",")
		for _, rule := range rules {
			ruleDetails := strings.Split(rule, "bag")
			ruleBag := strings.TrimSpace(ruleDetails[0])
			if ruleBag == "no other" {
				continue
			}
			ruleBagName := ruleBag[2:]
			if ruleBagName == "shiny gold" {
				shinyGoldHolders[bagName] = bagName
			}

			ruleBagNames = append(ruleBagNames, ruleBagName)
		}

		bagCollection[bagName] = ruleBagNames
	}
	totalShinyGolds := 0
	totalShinyGolds = checkBag("shiny gold", bagCollection, totalShinyGolds)

	fmt.Println(totalShinyGolds)
}

func checkBag(findBag string, bagCollection map[string][]string, counter int) int {
	for _, containedBags := range bagCollection {
		for _, bag := range containedBags {
			if bag == findBag {
				counter ++
				continue
			}
			counter = lookup(findBag, bagCollection[bag], counter)
		}
	}

	return counter
}

func lookup(findBag string, containedBags []string, counter int) int {
	for _, bag := range containedBags {
		if bag == findBag {
			counter ++
			continue
		}

		counter = lookup("shiny gold", containedBags, counter)
	}

	return counter
}
