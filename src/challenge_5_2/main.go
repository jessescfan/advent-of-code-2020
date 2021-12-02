package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const FileName = "input_5"

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
	seats := map[int]int{}

	maxId := 0
	for _, record := range lines {
		rows := record[0:7]
		columns := record[7:10]
		row := calcSpot(rows, 127, 0, 0, "F")
		column := calcSpot(columns, 7, 0, 0, "L")
		seat := (row * 8) + column
		if seat > maxId {
			maxId = seat
		}
		seats[seat] = seat
	}

	emptySeat := 0

	for i := 0; i < maxId; i++ {
		_, ok := seats[i]
		if !ok {
			emptySeat = i
		}
	}

	fmt.Println("empty seat: ", emptySeat)
}

func calcSpot(collection string, top int, mid int, bottom int, delimiter string) int  {
	for _, letter := range collection {
		mid = (top + bottom) / 2
		if string(letter) == delimiter {
			top = mid
			continue
		}

		bottom = mid
	}

	return top
}
