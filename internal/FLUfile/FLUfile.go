package FLUfile

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/DimKush/CopyFilesUtil/internal/InputParams"
	"github.com/cheggaaa/pb"
)

var flu_default_name = "FLUproc.csv"

func parseUnit(slc []string) (InputParams.Unit, error) {
	var unit InputParams.Unit
	unit.SetFromPathFile(slc[0])
	unit.SetToPathFile(slc[1])

	if slc[2] != "" {
		tmpV, err := strconv.Atoi(slc[2])
		if err != nil {
			err := errors.New("Offset parameter is not a number.")
			return InputParams.Unit{}, err
		}
		unit.SetOffset(tmpV)
	}
	if slc[3] != "" {
		tmpV, err := strconv.Atoi(slc[3])
		if err != nil {
			err := errors.New("Limit parameter is not a number.")
			return InputParams.Unit{}, err
		}
		unit.SetLimit(tmpV)
	}

	return unit, nil
}

func parseFile(path string) (units []InputParams.Unit, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4
	reader.Comment = '#'

	cnt := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if strings.ToLower(record[0]) == "from" && strings.ToLower(record[1]) == "to" {
			continue
		}

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		unit, err := parseUnit(record)
		if err != nil {
			return nil, err
		}

		units = append(units, unit)
		cnt++
	}

	return units, nil
}

func ProcessFLUfile(path string) error {
	if !strings.Contains(path, flu_default_name) {
		err := errors.New("path doesn't contain the path to FLUproc.csv file.")
		return err
	}
	units, err := parseFile(path)
	if err != nil {
		return err
	}
	if len(units) < 1 {
		err := errors.New("Empty FLUproc.csv file.")
		return err
	}
	fmt.Println(units)
	// parallel execution
	bar := pb.StartNew(len(units))

	var wg sync.WaitGroup
	for _, val := range units {

		wg.Add(1)
		tmp := val

		fmt.Print(tmp)
		go func(param InputParams.Unit) {
			defer wg.Done()
			//defer bar.Increment()
			//defer bar.Update()

			param.Process()
		}(tmp)
		bar.Increment()
		bar.Update()
	}

	bar.Update()
	wg.Wait()
	bar.Finish()

	return nil
	// try to open

}
