package csv

import (
	"encoding/csv"
	"io"
)

type Reader struct {
	reader  io.Reader
	headers []string
}

func NewReader(r io.Reader) *Reader {
	return &Reader{reader: r}
}

func (r *Reader) SetHeaders(headers []string) {
	r.headers = headers
}

func (r *Reader) Headers() []string {
	return r.headers
}

func (r *Reader) Range() chan *Result {
	resultChan := make(chan *Result)

	go func() {
		reader := csv.NewReader(r.reader)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				resultChan <- &Result{err: err}
			}
			resultChan <- &Result{row: r.createRow(record)}
		}
		close(resultChan)
	}()

	return resultChan
}

func (r *Reader) createRow(record []string) *Row {
	row := NewRow()
	numHeaders := len(r.headers)
	for i, value := range record {
		if i >= numHeaders {
			break
		}
		if r.headers[i] == "" {
			continue
		}
		row.AddValue(r.headers[i], value)
	}

	return row
}
