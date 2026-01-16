package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// version
var version = "v1.0.0, 2026-01-16\nhttps://github.com/cyclone-github/container_truncator"

func versionFunc() {
	fmt.Fprintln(os.Stderr, version)
}

func welcome() {
	fmt.Println(" ----------------------------------- ")
	fmt.Println("|   Cyclone's Container Truncator   |")
	fmt.Println("| Bestcrypt / Truecrypt / Veracrypt |")
	fmt.Println(" ----------------------------------- ")
	fmt.Println(version)
	fmt.Println()
	fmt.Println("Program will truncate a Bestcrypt, Truecrypt or Veracrypt Container")
	fmt.Println("and save a new file as 'truncated_orginal_filename'")
	fmt.Println("You can then run .tc & .vc files with hashcat")
	fmt.Println("Program only supports .jbc .tc & .vc container files")
	fmt.Println()
}

func truncate() {
	// get all supported container files in current dir
	extensions := []string{".jbc", ".tc", ".vc"}
	var files []string

	for _, ext := range extensions {
		matches, err := filepath.Glob("*" + ext)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		files = append(files, matches...)
	}

	if len(files) == 0 {
		fmt.Println("No files found with extension .jbc .tc or .vc")
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

	// print list of files and prompt user to select one
	fmt.Println("Select a container file to truncate:")
	for i, file := range truncatedFiles {
		fmt.Printf("%d) %s\n", i+1, file)
	}
	var selected int
	fmt.Scan(&selected)

	// get selected file
	inputFile := truncatedFiles[selected-1]

	// get output file name by prepending "_truncate" to input file name
	outputFile := "truncate_" + inputFile

	// check if output file already exists
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		fmt.Println("Error: Output file already exists.")
		return
	}

	// open input file
	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer in.Close()

	// open output file
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer out.Close()

	// determine header size based on container type
	var headerSize int64 = 512 // default for .tc / .vc
	if strings.HasSuffix(strings.ToLower(inputFile), ".jbc") {
		headerSize = 1536 // default for BestCrypt v8/9
	}

	// copy header bytes from input file to output file
	_, err = io.CopyN(out, in, headerSize)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File truncated successfully.")
	}
}

func main() {
	version := flag.Bool("version", false, "Program version:")
	cyclone := flag.Bool("cyclone", false, "dev info")
	help := flag.Bool("help", false, "Prints help:")
	flag.Parse()

	// run sanity checks for special flags
	if *version {
		versionFunc()
		os.Exit(0)
	}
	if *cyclone {
		decodedStr, err := base64.StdEncoding.DecodeString("Q29kZWQgYnkgY3ljbG9uZSA7KQo=")
		if err != nil {
			fmt.Fprintln(os.Stderr, "--> Cannot decode base64 string. <--")
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, string(decodedStr))
		os.Exit(0)
	}
	if *help {
		welcome()
		os.Exit(0)
	}

	welcome()
	truncate()
}

// end of program
