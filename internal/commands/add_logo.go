package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddLogo(videoPath, logoPath string) {
	folderName := "withLogo"
	if err := os.MkdirAll(folderName, 0755); err != nil {
		fmt.Printf("Error creating folder '%s': %v\n", folderName, err)
		os.Exit(1)
	}

	fileName := filepath.Base(videoPath)
	subfolderName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	if err := os.MkdirAll(filepath.Join(folderName, subfolderName), 0755); err != nil {
		fmt.Printf("Error creating subfolder '%s': %v\n", subfolderName, err)
		os.Exit(1)
	}

	outputPath := filepath.Join(folderName, subfolderName, fileName)

	fmt.Printf("Adding logo to '%s'...\n", videoPath)

	if err := executeCommand(
		"ffmpeg",
		"-i", videoPath,
		"-i", logoPath,
		"-filter_complex",
		"[1]scale=iw*1.3:ih*1.3,format=rgba,colorchannelmixer=aa=0.8[logo];[0][logo]overlay=1890:1047",
		"-c:v", "h264_nvenc",
		"-preset", "slow",
		"-cq", "18",
		"-profile:v", "high",
		"-pix_fmt", "yuv420p",
		"-c:a", "copy",
		"-movflags", "+faststart",
		"-y", outputPath,
	); err != nil {
		fmt.Printf("Error adding logo: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Logo added to: %s\n", outputPath)
}
