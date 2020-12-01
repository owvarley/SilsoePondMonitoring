package main

import (
	"github.com/owvarley/SilsoePondMonitoring/PondTotal.Go/sumupper"
	"flag"
)

func main() {
	dirPath := flag.String("dir", "./", "Directory to parse CSV files from")
	flag.Parse()

    sumupper.SumUp(*dirPath)
}