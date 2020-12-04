package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const FileName = "input_3"

type Slopes struct {
	Right int
	Down int
}

func main() {

	dir, err := filepath.Abs("src/inputs/" + FileName)
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	position := 1
	line := 1
	trees := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if position > len(text) {
			position = position - len(text)
		}
		character := text[position-1: position]

		if character == "#" {
			fmt.Println("Tree", text, line, position)
			trees++
		}
		position += 3
		line ++
	}

	fmt.Println(trees)
}
//60 66 77