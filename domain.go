package cookiedisk

import "strings"

type domain struct {
	httpOnlyPrefix string
	dot            string
	name           string
}

func newDomain(name string, tailMatch bool, httpOnly bool) (d domain) {
	d.name = name
	if tailMatch && !strings.HasPrefix(d.name, ".") {
		d.dot = "."
	}
	if httpOnly {
		d.httpOnlyPrefix = "#HttpOnly_"
	}
	return
}

func (d domain) String() string {
	return d.httpOnlyPrefix + d.dot + d.name
}
