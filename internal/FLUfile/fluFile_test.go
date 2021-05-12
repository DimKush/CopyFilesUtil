package FLUfile

import (
	"os"
	"testing"
)

func prepareContentFile() error {
	file, err := os.OpenFile("FLUproc.csv", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	str := "from,to,offset,limit\n" +
		"/home/dim/Desktop/Go/CopyFilesUtil/cmd/main.go,/home/dim/Desktop/trash/m.go,50,100\n" +
		"/home/dim/Desktop/Go/CopyFilesUtil/cmd/main.go,/home/dim/Desktop/trash/FLUfile.go,,\n"

	_, err = file.Write([]byte(str))
	if err != nil {
		return err
	}

	return nil
}

func TestParseUnit(t *testing.T) {
	err := prepareContentFile()

	if err != nil {
		t.Fatalf("Error in TestCheckFile(). CreateTestedFiles() returned an error : %s", err)
	}

	os.Remove("FLUproc.csv")
}
