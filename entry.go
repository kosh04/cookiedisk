// Package cookiedisk provides to use Netscape cookie file format (cookie.txt).
package cookiedisk

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Entry is Netscape cookie format
// see also: CURL/lib/cookie.c:get_netscape_format()
type Entry struct {
	Domain    domain
	TailMatch bool // weather we do tail-matchning of the domain name
	Path      string
	Secure    bool
	Expires   time.Time
	Name      string
	Value     string
}

// NewEntry creates Netscape cookie entry from http.Cookie
func NewEntry(u url.URL, c http.Cookie) *Entry {
	domain := c.Domain
	if domain == "" {
		domain = u.Host
	}

	tailMatch := false

	path := c.Path
	if path == "" {
		path = "/"
	}

	return &Entry{
		Domain:    newDomain(domain, tailMatch, c.HttpOnly),
		TailMatch: tailMatch,
		Path:      path,
		Secure:    c.Secure,
		Expires:   c.Expires,
		Name:      c.Name,
		Value:     c.Value,
	}
}

// record returns csv record converted from http.Cookie
func (e Entry) record() []string {
	var utime int64
	if !e.Expires.Equal(time.Time{}) {
		utime = e.Expires.Unix()
	}

	strBool := func(b bool) string {
		return strings.ToUpper(strconv.FormatBool(b))
	}

	return []string{
		e.Domain.String(),
		strBool(e.TailMatch),
		e.Path,
		strBool(e.Secure),
		strconv.FormatInt(utime, 10),
		e.Name,
		e.Value,
	}
}

func (e Entry) String() string {
	return strings.Join(e.record(), "\t")
}
