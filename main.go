package main

import (
	"flag"
	"os"

	"github.com/av-ugolkov/gopkg/logger"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "path to the file")

	flag.Parse()

	if path == "" {
		logger.Errorf("path is required")
		os.Exit(1)
	}

	var inst map[any]any
	err := generateStructure(path, inst)
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
}
