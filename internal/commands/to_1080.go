package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func To1080(inputPath string) {
	folderName := "1080"
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

	outputPath := filepath.Join(folderName, subfolderName, fileName)

	fmt.Printf("Converting '%s' to 1080...\n", inputPath)

	if err := executeCommand(
		"ffmpeg",
		"-i", inputPath,
		"-vf", "scale=1920:1080",
		"-c:v", "h264_nvenc",
		"-preset", "p7",
		"-tune", "hq",
		"-profile:v", "high",
		"-rc", "vbr",
		"-cq:v", "16",
		"-b:v", "10M",
		"-maxrate", "14M",
		"-bufsize", "28M",
		"-spatial-aq", "1",
		"-temporal-aq", "1",
		"-rc-lookahead", "32",
		"-surfaces", "64",
		"-pix_fmt", "yuv420p",
		"-c:a", "aac",
		"-b:a", "320k",
		"-movflags", "+faststart",
		"-y", outputPath,
	); err != nil {
		fmt.Printf("Error converting to 1080: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Conversion completed: %s\n", outputPath)
}
