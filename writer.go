package cookiedisk

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type writer struct {
	w *csv.Writer
}

func newWriter(w io.Writer) *writer {
	cw := csv.NewWriter(w)
	cw.Comma = '\t'
	return &writer{cw}
}

func (w *writer) WriteAll(u *url.URL, cookies []*http.Cookie) error {
	var records [][]string
	for _, c := range cookies {
		fmt.Printf("Before:%#v\n", c)
		e := NewEntry(*u, *c)
		fmt.Printf("After: %v\n", e.String())
		records = append(records, e.record())
	}
	return w.w.WriteAll(records)
}

// WriteFile write cookies to filename
func WriteFile(filename string, u *url.URL, cookies []*http.Cookie) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(f, "# This file was generated! Edit at your own risk.")

	w := newWriter(f)
	return w.WriteAll(u, cookies)
}
