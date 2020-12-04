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
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	totalTrees := 1
	slopes := []Slopes{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	}

	for _, slope := range slopes {
		trees := 0
		position := 0
		lineWidth := len(lines[0])

//borrowed this from another solution
		for i := slope.Down; i < len(lines); i += slope.Down {
			position = (position + slope.Right) % lineWidth
			if string(lines[i][position]) == "#" {
				trees++
			}
		}
		totalTrees = trees * totalTrees
	}
	fmt.Println("totalTrees", totalTrees)
}
//3737923200