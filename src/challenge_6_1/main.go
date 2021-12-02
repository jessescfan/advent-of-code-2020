package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const FileName = "input_6"

func main() {
	dir, err := filepath.Abs("src/inputs/" + FileName)
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	var groupsText string
	for scanner.Scan() {
		if scanner.Text() == "" {
			lines = append(lines, groupsText)
			groupsText = ""
		} else {
			groupsText += scanner.Text() + " "
		}
	}
	lines = append(lines, groupsText)

	totalCounts := 0
	for _, group := range lines {
		groupAnswers := strings.Split(strings.TrimSpace(group), " ")
		answerMap := make(map[string]string)
		for _, personAnswers := range groupAnswers {
			for _, letter := range personAnswers {
				answerMap[string(letter)] = string(letter)
			}
		}
		totalCounts += len(answerMap)
	}

	fmt.Println("Count ", totalCounts)
}
