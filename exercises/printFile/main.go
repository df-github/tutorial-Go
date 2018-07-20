package main

import (
	"os"
	"fmt"
	"io"
	"path/filepath"
)

func printFileContent(filename string) {
	file := getFile(filename)
	io.Copy(os.Stdout, file)
}

func getFile(filename string) *os.File {
	filePath := getAbsolutePath(filename)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Could not open file", filename)
		os.Exit(1)
	}
	return file
}

func getAbsolutePath(filename string) string {
	filePath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("Could not get the abasolute path of the given filename", filename)
		os.Exit(1)
	}
	return filePath
}

func main() {
	arguments := os.Args
	if len(arguments) > 1 {
		filename := arguments[1]
		printFileContent(filename)
	} else {
		fmt.Println("filename must be given for running this programm")
		fmt.Println("run main.go <filename>")
	}
}
