package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Jibaru/iwa/internal/commands"
)

func main() {
	inputFile := flag.String("i", "", "Input file path")
	logoFile := flag.String("l", "", "Logo file path")
	commandName := flag.String("c", "", "Command: avi2mp4, 2k, 1080")
	flag.Parse()

	if *inputFile == "" || *commandName == "" {
		fmt.Println("Usage: app -i <input_file> -c <avi2mp4|2k|1080>")
		os.Exit(1)
	}

	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' does not exist\n", *inputFile)
		os.Exit(1)
	}

	switch strings.ToLower(*commandName) {
	case "avi2mp4":
		commands.AVIToMP4(*inputFile)
	case "2k":
		commands.To2K(*inputFile)
	case "1080":
		commands.To1080(*inputFile)
	case "addlogo":
		if *logoFile == "" {
			fmt.Println("Usage: app -i <input_file> -c addlogo -l <logo_file>")
			os.Exit(1)
		}

		commands.AddLogo(*inputFile, *logoFile)
	default:
		fmt.Printf("Error: Unknown conversion type '%s'. Use: mp4, 2k, or 1080\n", *commandName)
		os.Exit(1)
	}
}
