package main

import (
	"flag"
)

// flag params
var fromPath string
var toPath string
var limit int
var offset int
var procCsvPath string

func init() {
	flag.StringVar(&procCsvPath, "flu", "./FLUproc.csv", "path to the processing file")
	flag.StringVar(&fromPath, "from", "", "the file to read from")
	flag.StringVar(&toPath, "to", "", "the file to write")
	flag.IntVar(&limit, "limit", 0, "limit of copied bytes")
	flag.IntVar(&offset, "offset", 0, "offset of copied bytes")

	flag.Parse()
}

func main() {
	if procCsvPath != "" {
		ProcessFLUfile(procCsvPath)
	} else {

	}
}
