package cookiedisk

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var testCases = []struct {
	text    string
	cookies []*http.Cookie
}{
	{
		text:    ``,
		cookies: []*http.Cookie{},
	},
	// Irreversible conversion
	// {
	// 	text: `# Comment`,
	// 	cookie: []*http.Cookie{},
	// },
	{
		text: `httpbin.org	FALSE	/	TRUE	0	k1	v1`,
		cookies: []*http.Cookie{
			{
				Domain: "httpbin.org",
				Path:   "/",
				Secure: true,
				Name:   "k1",
				Value:  "v1",
			},
		},
	},
	{
		text: `httpbin.org	FALSE	/	FALSE	0	k2	v2`,
		cookies: []*http.Cookie{
			{
				Domain: "httpbin.org",
				Path:   "/",
				Secure: false,
				Name:   "k2",
				Value:  "v2",
			},
		},
	},
}

func TestRead(t *testing.T) {
	for _, tc := range testCases {
		cookies := Read(strings.NewReader(tc.text))
		if len(cookies) != len(tc.cookies) {
			t.Errorf("length %v, want %v", len(cookies), len(tc.cookies))
		}
		if !reflect.DeepEqual(cookies, tc.cookies) {
			t.Errorf("cookies %v, want %v", cookies, tc.cookies)
		}
	}
}

func TestReadFile(t *testing.T) {
	cases := []struct {
		filename string
		want     []*http.Cookie
	}{
		{
			filename: "data/cookie.txt",
			want: []*http.Cookie{
				{
					Domain: "httpbin.org",
					Path:   "/",
					Secure: false,
					Name:   "k1",
					Value:  "v1",
				},
				{
					Domain: "httpbin.org",
					Path:   "/",
					Secure: true,
					Name:   "k2",
					Value:  "v2",
				},
			},
		},
	}
	for _, tc := range cases {
		cookies, err := ReadFile(tc.filename)
		if err != nil {
			t.Errorf("%s %v\n", tc.filename, err)
		}
		if !reflect.DeepEqual(cookies, tc.want) {
			t.Errorf("%v but want \n%v\n", cookies, tc.want)
		}

	}
}
