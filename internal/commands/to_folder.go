package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func ToFolder(inputPath string, parentFolderName string) {
	folderName := parentFolderName
	if err := os.MkdirAll(folderName, 0755); err != nil {
		fmt.Printf("Error creating folder '%s': %v\n", folderName, err)
		os.Exit(1)
	}

	fileName := filepath.Base(inputPath)
	subfolderName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	if err := os.MkdirAll(filepath.Join(folderName, subfolderName), 0755); err != nil {
		fmt.Printf("Error creating subfolder '%s': %v\n", subfolderName, err)
		os.Exit(1)
	}

	// Move the input file to the subfolder
	destPath := filepath.Join(folderName, subfolderName, fileName)
	if err := os.Rename(inputPath, destPath); err != nil {
		fmt.Printf("Error moving file '%s' to '%s': %v\n", inputPath, destPath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully moved '%s' to '%s'\n", inputPath, destPath)
}
