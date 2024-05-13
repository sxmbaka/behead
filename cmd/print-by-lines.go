package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		if err.Error() == "EOF" {
			return
		}
		if os.IsNotExist(err) {
			fmt.Println("No such file or directory!")
			os.Exit(1)
		}
		panic(err)
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
		check(err)
		if numbered {
			fmt.Printf("%6d | %s", i+1, line)

		} else {
			fmt.Printf("%s", line)
		}
	}
}
