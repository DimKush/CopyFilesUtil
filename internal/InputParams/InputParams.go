package InputParams

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Unit struct {
	from   string
	to     string
	limit  int
	offset int
}

func (d *Unit) SetFromPathFile(fromPathFile string) {
	d.from = fromPathFile
}

func (d *Unit) SetToPathFile(toPathFile string) {
	d.to = toPathFile
}

func (d *Unit) SetLimit(limit int) {
	d.limit = limit
}

func (d *Unit) SetOffset(offset int) {
	d.offset = offset
}

func (d *Unit) GetLimit() int {
	return d.limit
}

func (d *Unit) GetOffset() int {
	return d.offset
}

func (d *Unit) Process() error {
	fin, err := os.Open(d.from)

	if err != nil {
		fmt.Println(err)
		return err
	}

	br := bufio.NewReader(fin)

	buf, err := ioutil.ReadAll(br)

	if d.limit == 0 {
		d.limit = len(buf)
	}

	if len(buf) < d.limit {
		return errors.New(fmt.Sprintf("Incorrect vallues. limit %d can't be higher that file length %d", d.limit, len(buf)))
	}

	fout, err := os.OpenFile(d.to, os.O_CREATE|os.O_WRONLY, 0666)
	bufout, err := fout.Write(buf[d.offset:d.limit])
	fmt.Printf("%d", bufout)

	return nil
}
