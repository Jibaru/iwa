package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Jibaru/iwa/internal/commands"
)

type stringSlice []string

func (s *stringSlice) String() string {
	return strings.Join(*s, ", ")
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var inputFiles stringSlice

	logoFile := flag.String("l", "", "Logo file path")
	commandName := flag.String("c", "", "Command: avi2mp4, 2k, 1080, addlogo, to4kfolder, to2kfolder, to1080folder")

	flag.Var(&inputFiles, "i", "Input file path (can be repeated)")
	flag.Parse()

	if len(inputFiles) == 0 || *commandName == "" {
		fmt.Println("Usage: app -i <file1> -i <file2> -c <avi2mp4|2k|1080|addlogo|to4kfolder|to2kfolder|to1080folder>")
		os.Exit(1)
	}

	for _, f := range inputFiles {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			fmt.Printf("Error: File '%s' does not exist\n", f)
			os.Exit(1)
		}
	}

	for _, input := range inputFiles {
		switch strings.ToLower(*commandName) {
		case "avi2mp4":
			commands.AVIToMP4(input)
		case "2k":
			commands.To2K(input)
		case "1080":
			commands.To1080(input)
		case "addlogo":
			if *logoFile == "" {
				fmt.Println("Usage: app -i <file> -c addlogo -l <logo_file>")
				os.Exit(1)
			}
			commands.AddLogo(input, *logoFile)
		case "to4kfolder":
			commands.ToFolder(input, "4K")
		case "to2kfolder":
			commands.ToFolder(input, "2K")
		case "to1080folder":
			commands.ToFolder(input, "1080")
		default:
			fmt.Printf("Error: Unknown conversion type '%s'\n", *commandName)
			os.Exit(1)
		}
	}
}
