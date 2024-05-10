package cmd

import (
	"bufio"
	"log"
	"os"
)

func countLines(path string) int {
	path = resolvePath(path)
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lineCount
}

func countBytes(path string) int {
	path = resolvePath(path)
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	fileInfo, err := file.Stat()
	check(err)
	byteCount := fileInfo.Size()
	return int(byteCount)
}
