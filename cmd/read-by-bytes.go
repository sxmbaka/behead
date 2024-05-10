package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func printByBytes(path string, bytes int, printHeader bool, numbered bool) {
	path = resolvePath(path)
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	if printHeader {
		fmt.Printf("======> %s <======\n", path)
	}
	for i := 0; i < bytes; i++ {
		char, err := reader.ReadByte()
		if err != nil {
			// Check if the error is due to EOF
			if err.Error() == "EOF" {
				// Reached end of file, break out of the loop
				break
			} else {
				// Other error occurred
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}
		}
		if numbered {
			fmt.Printf("%6d | %c", i+1, char)
		}
		fmt.Printf("%c", char)
	}
}
