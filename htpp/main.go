package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
}
