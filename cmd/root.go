package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "behead [FILE]",
	Short: "A simple CLI that resembles the GNU's head command.",
	Long:  `A simple command-line application that resembles the GNU's head command. It reads the first few (10 by default) lines of a file and writes them to standard output. It is useful for examining the beginning of a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var printHeader bool = false
		var delim byte = '\n'

		lines, _ := cmd.Flags().GetInt("lines")
		bytes, _ := cmd.Flags().GetInt("bytes")
		silent := cmd.Flags().Changed("silent")
		verbose := cmd.Flags().Changed("verbose")
		zeroTerminated := cmd.Flags().Changed("zero-terminated")
		params := cmd.Flags().Changed("params")
		numbered := cmd.Flags().Changed("numbered")
		inverted := cmd.Flags().Changed("inverted")

		if zeroTerminated {
			delim = '\x00'
		}

		if verbose || !silent && len(args) > 1 {
			printHeader = true
		}

		if len(args) == 0 {
			// Read from standard input (os.Stdin) when no file arguments are provided
			cmd.Println("No files provided!\nSee behead --help for usage.")
			os.Exit(1)
		}

		if params {
			checkParams(lines, bytes, silent, verbose, zeroTerminated, numbered, args)
			os.Exit(0)
		}

		if bytes == 0 {
			for _, path := range args {
				if inverted {
					lines = countLines(path) - lines
				}
				printByLines(path, lines, delim, printHeader, numbered)
			}
		} else {
			for _, path := range args {
				if inverted {
					bytes = countBytes(path) - bytes
				}
				printByBytes(path, lines, printHeader, numbered)
			}
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().IntP("lines", "n", 10, "decides the number of lines to print out.")
	rootCmd.Flags().IntP("bytes", "b", 0, "decides the number of bytes to print out.")
	rootCmd.Flags().BoolP("silent", "s", false, "never print file headers.")
	rootCmd.Flags().BoolP("verbose", "v", false, "always print file headers.")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimeter is NUL (\\0), not newline (\\n).")
	rootCmd.Flags().BoolP("params", "p", false, "check tool parameters for testing.")
	rootCmd.Flags().Bool("numbered", false, "show line numbers")
	rootCmd.Flags().BoolP("inverted", "i", false, "when used, the result is the file\nexcluding the last NUM lines or bytes.")
}
