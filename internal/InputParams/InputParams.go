package InputParams

import "github.com/DimKush/CopyFilesUtil/internal/CopyProcess"

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
	err := CopyProcess.CopyFile(d.from, d.to, d.offset, d.limit)
	if err != nil {
		return err
	}
	return nil
}
