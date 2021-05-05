package InputParams

type Unit struct {
	fromPath string
	toPath   string
	limit    int
	offset   int
}

func (d *Unit) SetFromPath(fromPath string) {
	d.fromPath = fromPath
}

func (d *Unit) SetToPath(toPath string) {
	d.toPath = toPath
}

func (d *Unit) SetLimit(limit int) {
	d.limit = limit
}

func (d *Unit) SetOffset(offset int) {
	d.offset = offset
}
