package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const FileName = "input_2"

type Record struct {
	letter string
	min int
	max int
	password string
}

func main() {
	dict := input("input_2")
	valid := 0

	for _, pass := range dict {
		first := pass.password[pass.min-1: pass.min]
		second := pass.password[pass.max-1: pass.max]
		if first != pass.letter && second != pass.letter {
			continue
		}
		if first == second {
			continue
		}

		valid++
	}

	fmt.Println(valid)
}

func input(fileName string) map[string]Record {
	n := map[string]Record{}
	dir, err := filepath.Abs("src/inputs/" + fileName)
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		minMax := strings.Split(text[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		n[scanner.Text()] = Record{
			letter: text[1][0:1],
			min: min,
			max: max,
			password: text[2],
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return n
}