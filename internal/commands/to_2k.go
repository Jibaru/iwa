package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func To2K(inputPath string) {
	folderName := "2K"
	if err := os.MkdirAll(folderName, 0755); err != nil {
		fmt.Printf("Error creating folder '%s': %v\n", folderName, err)
		os.Exit(1)
	}

	fileName := filepath.Base(inputPath)
	outputPath := filepath.Join(folderName, fileName)

	fmt.Printf("Converting '%s' to 2K...\n", inputPath)

	command := fmt.Sprintf("ffmpeg -i \"%s\" -vf scale=2560:1440 -c:v h264_nvenc -preset p7 -tune hq -profile:v high -rc vbr -cq:v 16 -b:v 16M -maxrate 22M -bufsize 44M -spatial-aq 1 -temporal-aq 1 -rc-lookahead 32 -surfaces 64 -pix_fmt yuv420p -c:a aac -b:a 320k -movflags +faststart -y \"%s\"", inputPath, outputPath)

	if err := executeCommand(command); err != nil {
		fmt.Printf("Error converting to 2K: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Conversion completed: %s\n", outputPath)
}
