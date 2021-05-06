package FLUfile

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DimKush/CopyFilesUtil/internal/InputParams"
)

var flu_default_name = "FLUproc.csv"

func parseUnit(slc []string) (InputParams.Unit, error) {
	fmt.Println(slc)
	var unit InputParams.Unit
	unit.SetFromPath(slc[0])
	unit.SetToPath(slc[1])

	tmpV, err := strconv.Atoi(slc[2])
	if err != nil {
		err := errors.New("Offset parameter is not a number.")
		return InputParams.Unit{}, err
	}
	unit.SetOffset(tmpV)

	tmpV, err = strconv.Atoi(slc[3])
	if err != nil {
		err := errors.New("Limit parameter is not a number.")
		return InputParams.Unit{}, err
	}
	unit.SetOffset(tmpV)

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
		if cnt == 0 {
			continue
		}
		fmt.Println("1")
		record, err := reader.Read()
		fmt.Println("2")
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
}

//func

func ProcessFLUfile(path string) error {
	if !strings.Contains(path, flu_default_name) {
		err := errors.New("path doesn't contain the path to FLUproc.csv file.")
		return err
	}
	_, err := parseFile(path)
	if err != nil {
		return err
	}

	return nil
	// try to open

}
