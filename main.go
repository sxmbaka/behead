package main

import (
	"fmt"

	"github.com/sxmbaka/behead/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		// Handle error
		fmt.Println(err)
	}
}
