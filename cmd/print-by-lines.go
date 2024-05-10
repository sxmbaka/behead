package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		// panic(err)
		fmt.Println("No such file or directory!")
		os.Exit(1)
	}
}

func printByLines(path string, lines int, delim byte, printHeader bool, numbered bool) {
	path = resolvePath(path)
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	if printHeader {
		fmt.Printf("======> %s <======\n", path)
	}
	for i := 0; i < lines; i++ {
		line, err := reader.ReadString(delim)
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
			fmt.Printf("%6d | %s", i+1, line)
			continue
		}
		fmt.Printf("%s", line)
	}
}
