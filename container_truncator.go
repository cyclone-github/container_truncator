package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// version history
var version = "v0.1.0, 2023-01-23; initial release"

// clear screen function
func clearScreen() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Note, program has not been tested on Mac OS")
		time.Sleep(5 * time.Second)
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func welcome() {
	fmt.Println(" ----------------------- ")
	fmt.Println("|       Cyclone's       |")
	fmt.Println("| Truecrypt / Veracrypt |")
	fmt.Println("|  Container Truncater  |")
	fmt.Println(" ----------------------- ")
	fmt.Println(version)
	fmt.Println()
	fmt.Println("Program will truncate a Truecrypt or Veracrypt Container")
	fmt.Println("and save a new file as 'truncated_orginal_filename'.")
	fmt.Println("You can then run this file with hashcat.")
	fmt.Println("Program only supports .tc & .vc container files.")
	fmt.Println()
}

func truncate() {
	// Get all files in current directory with .tc or .vc extension
	files, err := filepath.Glob("*.tc")
	files2, err := filepath.Glob("*.vc")
	files = append(files, files2...)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(files) == 0 {
		fmt.Println("No files found with extension .tc or .vc")
		return
	}

	// filter files that have already been truncated
	var truncatedFiles []string
	for _, file := range files {
		if !strings.HasPrefix(file, "truncate_") {
			truncatedFiles = append(truncatedFiles, file)
		}
	}
	if len(truncatedFiles) == 0 {
		fmt.Println("All files have already been truncated.")
		return
	}

	// Print list of files and ask user to select one
	fmt.Println("Select a container file to truncate:")
	for i, file := range truncatedFiles {
		fmt.Printf("%d) %s\n", i+1, file)
	}
	var selected int
	fmt.Scan(&selected)

	// Get selected file
	inputFile := truncatedFiles[selected-1]

	// Get output file name by prepending "_truncate" to input file name
	outputFile := "truncate_" + inputFile

	// Check if output file already exists
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		fmt.Println("Error: Output file already exists.")
		return
	}

	// Open input file
	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer in.Close()

	// Open output file
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer out.Close()

	// Copy 512 bytes from input file to output file
	_, err = io.CopyN(out, in, 512)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File truncated successfully.")
	}
}

func main() {
	clearScreen()
	welcome()
	truncate()
}

// end of program
