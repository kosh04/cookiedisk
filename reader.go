package cookiedisk

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// cookie.txt reader (csv wrapper)
type reader struct {
	r *csv.Reader
}

var (
	errMissingRecord = errors.New("missing cookie record")
)

func newReader(r io.Reader) *reader {
	cr := csv.NewReader(r)
	cr.Comma = '\t'
	cr.Comment = '#'
	return &reader{cr}
}

func convert(record []string) (*http.Cookie, error) {
	if len(record) != 7 {
		return nil, errMissingRecord
	}
	domain, _, path, _secure, _expires, name, value := record[0], record[1], record[2], record[3], record[4], record[5], record[6]
	cookie := &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   path,
		Domain: domain,
	}

	secure, err := strconv.ParseBool(_secure)
	if err == nil {
		cookie.Secure = secure
	}

	expires, err := strconv.ParseInt(_expires, 10, 64)
	if err != nil {
		return nil, err
	}
	exptime := time.Unix(expires, 0)
	// zero probably means that it never expires,
	// or that it is good for as long as this session lasts.
	if expires > 0 {
		cookie.Expires = exptime
	}

	return cookie, nil
}

// Note: unreadable cookie record will be discard
func (r *reader) readAll() []*http.Cookie {
	cookies := []*http.Cookie{}
	for {
		record, err := r.r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}
		cookie, err := convert(record)
		if err != nil {
			continue
		}
		cookies = append(cookies, cookie)
	}
	return cookies
}

// Read cookies from stream
func Read(r io.Reader) []*http.Cookie {
	rr := newReader(r)
	return rr.readAll()
}

// ReadFile read cookies from file
func ReadFile(filename string) ([]*http.Cookie, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Read(f), nil
}
