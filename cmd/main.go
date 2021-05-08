package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DimKush/CopyFilesUtil/internal/CopyProcess"
	"github.com/DimKush/CopyFilesUtil/internal/FLUfile"
)

// flag params
var fromPath string
var toPath string
var limit int
var offset int
var procCsvPath string

func init() {
	flag.StringVar(&fromPath, "from", "", "the file to read from")
	flag.StringVar(&toPath, "to", "", "the file to write")
	flag.IntVar(&offset, "offset", 0, "offset of copied bytes")
	flag.IntVar(&limit, "limit", 0, "limit of copied bytes")

	flag.StringVar(&procCsvPath, "flu", "", "path to the processing file")
	flag.Parse()
}

func main() {
	var err error
	if procCsvPath != "" {
		err = FLUfile.ProcessFLUfile(procCsvPath)
	} else {
		err = CopyProcess.CopyFile(fromPath, toPath, offset, limit)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
