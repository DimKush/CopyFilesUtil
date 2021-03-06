package CopyProcess

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkOffset(offset int, limit int) error {
	if offset >= limit {
		return fmt.Errorf("Error offset %d can't be higher that limit %d.", offset, limit)
	}

	return nil
}

func CopyFile(from string, to string, offset int, limit int) error {
	fin, err := os.Open(from)

	if err != nil {
		fmt.Println(err)
		return err
	}

	br := bufio.NewReader(fin)

	buf, err := ioutil.ReadAll(br)
	if err != nil {
		return err
	}

	if limit == 0 {
		limit = len(buf)
	}

	if len(buf) < limit {
		return fmt.Errorf("Incorrect vallues. limit %d can't be higher that file length %d", limit, len(buf))
	}

	err = checkOffset(offset, limit)
	if err != nil {
		return err
	}

	fout, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	bufout, err := fout.Write(buf[offset:limit])
	if err != nil {
		return err
	}
	fmt.Printf("%d", bufout)

	return nil
}
