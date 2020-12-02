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

	for _, i2 := range dict {
		for _, i3 := range dict {

			if i3 + i2 > 2020 {
				continue
			}
			f := 2020 - (i2 + i3)

			if val, ok := dict[f]; ok {
				solution := i3 * val * i2
				fmt.Println(solution)
				return
			}
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