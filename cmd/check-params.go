package cmd

import "fmt"

func checkParams(lines int, bytes int, silent bool, verbose bool, zeroTerminated bool, numbered bool, args []string) {
	fmt.Printf("lines: %v\n", lines)
	fmt.Printf("bytes: %v\n", bytes)
	fmt.Printf("silent: %v\n", silent)
	fmt.Printf("verbose: %v\n", verbose)
	fmt.Printf("zeroTerminated: %v\n", zeroTerminated)
	fmt.Printf("numbered: %v\n", numbered)
	fmt.Printf("args: %v\n", args)
}
