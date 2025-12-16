package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AVIToMP4(inputPath string) {
	ext := filepath.Ext(inputPath)
	outputPath := strings.TrimSuffix(inputPath, ext) + ".mp4"

	fmt.Printf("Converting '%s' to MP4...\n", inputPath)

	if err := executeCommand(
		"ffmpeg",
		"-i", inputPath,
		"-c:v", "h264_nvenc",
		"-preset", "p7",
		"-tune", "hq",
		"-profile:v", "high",
		"-rc", "vbr",
		"-cq:v", "16",
		"-b:v", "25M",
		"-maxrate", "35M",
		"-bufsize", "70M",
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
		fmt.Printf("Error converting to MP4: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Conversion completed: %s\n", outputPath)
}
