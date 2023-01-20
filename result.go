package csv

type Result struct {
	row *Row
	err error
}

func (r *Result) Row() *Row {
	return r.row
}

func (r *Result) Error() error {
	return r.err
}
