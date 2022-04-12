package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := getFilename()
	file := openFile(filename)

	io.Copy(os.Stdout, file)
}

func getFilename() string {
	return os.Args[1]
}

func openFile(fn string) *os.File {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	return file
}
