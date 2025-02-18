package main

import (
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"time"

	"github.com/kbinani/screenshot"
)

func captureScreen(filename string) error {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		return fmt.Errorf("No active displays found")
	}

	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func startRecording(duration time.Duration) {
	folder := "recordings"
	os.MkdirAll(folder, os.ModePerm)

	fmt.Println("Recording started...")
	start := time.Now()
	frames := []string{}

	for time.Since(start) < duration {
		filename := fmt.Sprintf("%s/frame_%d.png", folder, time.Now().UnixNano())
		err := captureScreen(filename)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		frames = append(frames, filename)
		time.Sleep(time.Second / 60) // 60 FPS
	}

	fmt.Println("Recording finished!")
	convertToVideo(frames)
}

func convertToVideo(frames []string) {
	fmt.Println("Converting to video...")

	// Buat daftar gambar untuk FFmpeg
	listFile := "recordings/filelist.txt"
	list, err := os.Create(listFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer list.Close()

	for _, frame := range frames {
		fmt.Fprintf(list, "file '%s'\n", frame)
	}

	outputVideo := "recordings/output.mp4"
	cmd := exec.Command("ffmpeg", "-y", "-f", "concat", "-safe", "0", "-i", listFile, "-vf", "fps=60", "-pix_fmt", "yuv420p", outputVideo)
	err = cmd.Run()
	if err != nil {
		fmt.Println("FFmpeg Error:", err)
	} else {
		fmt.Println("Video saved as:", outputVideo)
	}

	// Hapus file gambar setelah konversi selesai
	for _, frame := range frames {
		os.Remove(frame)
	}
	os.Remove(listFile)
}

func main() {
	fmt.Println("Starting screen recording...")
	startRecording(10 * time.Second)
	fmt.Println("Program finished!")
}
