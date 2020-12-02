package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const FileName = "input_1"

func main() {
	dict := input(FileName)

	count := 0
	for i, i2 := range dict {
		f := 2020 - i2
		count ++
		if val, ok := dict[f]; ok {
			solution := dict[i] * val
			fmt.Println(solution)
			return
		}
	}
}


func input(fileName string) map[int]int {
	n := make (map[int]int)
	dir, err := filepath.Abs("src/inputs/" + fileName)
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())

		n[i] = i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return n
}