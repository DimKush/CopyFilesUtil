package CopyProcess

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func CopyFile(from string, to string, offset int, limit int) error {
	fin, err := os.Open(from)

	if err != nil {
		fmt.Println(err)
		return err
	}

	br := bufio.NewReader(fin)

	buf, err := ioutil.ReadAll(br)

	if limit == 0 {
		limit = len(buf)
	}

	if len(buf) < limit {
		return errors.New(fmt.Sprintf("Incorrect vallues. limit %d can't be higher that file length %d", limit, len(buf)))
	}

	fout, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY, 0666)
	bufout, err := fout.Write(buf[offset:limit])
	fmt.Printf("%d", bufout)

	return nil
}
