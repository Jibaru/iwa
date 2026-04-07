package commands

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	iwaraTargetMB    = 599.0
	iwaraMaxrateKbps = 16100
)

func Iwara(inputPath string) {
	folderName := "iwara"
	if err := os.MkdirAll(folderName, 0755); err != nil {
		fmt.Printf("Error creating folder '%s': %v\n", folderName, err)
		os.Exit(1)
	}

	fileName := filepath.Base(inputPath)
	baseName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	subfolderPath := filepath.Join(folderName, baseName)
	if err := os.MkdirAll(subfolderPath, 0755); err != nil {
		fmt.Printf("Error creating subfolder '%s': %v\n", baseName, err)
		os.Exit(1)
	}

	outputPath := filepath.Join(subfolderPath, baseName+".mp4")

	duration, err := probeDurationSeconds(inputPath)
	if err != nil {
		fmt.Printf("Error detecting duration for '%s': %v\n", inputPath, err)
		os.Exit(1)
	}
	if duration <= 0 {
		fmt.Printf("Error: invalid duration %.2fs for '%s'\n", duration, inputPath)
		os.Exit(1)
	}

	// bitrate_kbps = (target_MB * 8192) / duration_seconds
	videoBitrateKbps := int(math.Round((iwaraTargetMB * 8192.0) / duration))
	if videoBitrateKbps > iwaraMaxrateKbps {
		videoBitrateKbps = iwaraMaxrateKbps
	}

	bvArg := fmt.Sprintf("%dk", videoBitrateKbps)
	maxrateArg := fmt.Sprintf("%dk", iwaraMaxrateKbps)
	bufsizeArg := fmt.Sprintf("%dk", iwaraMaxrateKbps*2)

	fmt.Printf("Encoding '%s' for iwara (duration=%.2fs, target=%.0fMB, bitrate=%s, maxrate=%s)...\n",
		inputPath, duration, iwaraTargetMB, bvArg, maxrateArg)

	fmt.Println("Pass 1/2 (analysis)...")
	if err := executeCommand(
		"ffmpeg",
		"-y",
		"-i", inputPath,
		"-c:v", "h264_nvenc",
		"-b:v", bvArg,
		"-maxrate", maxrateArg,
		"-bufsize", bufsizeArg,
		"-rc", "vbr_hq",
		"-2pass", "1",
		"-preset", "p5",
		"-profile:v", "high",
		"-an",
		"-f", "mp4",
		os.DevNull,
	); err != nil {
		fmt.Printf("Error in pass 1: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Pass 2/2 (final)...")
	if err := executeCommand(
		"ffmpeg",
		"-y",
		"-i", inputPath,
		"-c:v", "h264_nvenc",
		"-b:v", bvArg,
		"-maxrate", maxrateArg,
		"-bufsize", bufsizeArg,
		"-rc", "vbr_hq",
		"-2pass", "1",
		"-preset", "p5",
		"-profile:v", "high",
		"-c:a", "aac",
		"-b:a", "128k",
		"-movflags", "+faststart",
		outputPath,
	); err != nil {
		fmt.Printf("Error in pass 2: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Conversion completed: %s\n", outputPath)
}

func probeDurationSeconds(inputPath string) (float64, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		inputPath,
	)
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
}
