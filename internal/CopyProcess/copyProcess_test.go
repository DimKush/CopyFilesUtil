package CopyProcess

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

//func helpers

func CreateTestedFiles() error {

	file, err := os.OpenFile("ut_in.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	str := "Hello unit test TestCheckOffset"
	_, err = file.Write([]byte(str))
	if err != nil {
		return err
	}

	return nil
}

func TestCheckOffset(t *testing.T) {
	err := checkOffset(10, 220)
	if err != nil {
		t.Fatalf("Error in TestCheckOffset. checkOffset() returned an error : %s", err)
	}
}

func TestCheckFilePositive(t *testing.T) {
	err := CreateTestedFiles()

	if err != nil {
		t.Fatalf("Error in TestCheckFile(). CreateTestedFiles() returned an error : %s", err)
	}

	err = CopyFile("ut_in.txt", "ut_out.txt", 0, 0)
	if err != nil {
		t.Fatalf("Error in TestCheckFile(). CreateTestedFiles() returned an error : %s", err)
	}

	fout, err := os.OpenFile("ut_out.txt", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatalf("Error in TestCheckFile(). %s", err)
	}

	buf, err := ioutil.ReadAll(fout)
	if err != nil {
		t.Fatalf("Error in TestCheckFile(). %s", err)
	}

	strout := string(buf)
	stroutExpect := "Hello unit test TestCheckOffset"

	if strout != stroutExpect {
		t.Fatalf("Error in TestCheckFile(). Unexpected string %s. Expected %s", strout, stroutExpect)
	}

	os.Remove("ut_in.txt")
	os.Remove("ut_out.txt")
}

func TestCheckFileNegative(t *testing.T) {
	err := CreateTestedFiles()

	if err != nil {
		t.Fatalf("Error in TestCheckFile(). CreateTestedFiles() returned an error : %s", err)
	}

	// check offset
	err = CopyFile("ut_in.txt", "ut_out.txt", 1000, 0)
	if err == nil {
		t.Fatalf("TestCheckFileNegative. Error expected.")
	}
	fmt.Println(err)

	//check limit
	err = CopyFile("ut_in.txt", "ut_out.txt", 0, 100000)
	if err == nil {
		t.Fatalf("TestCheckFileNegative. Error expected.")
	}
	fmt.Println(err)

	os.Remove("ut_in.txt")
	os.Remove("ut_out.txt")
}
