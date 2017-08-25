CookieDisk
==========

cookiedisk is Go package to use Netscape cookie file format (e.g, `cookie.txt`).


## Example

	import(
		"net/http"
		"net/http/cookiejar"
		"net/http/url"

		"github.com/kosh04/cookiedisk"
	)
	
	rawURL := "http://httpbin.org/cookies"
	cookies, _ := cookiedisk.ReadFile("./data/cookie.txt")
	u, _ := url.Parse(rawURL)
	jar, _ := cookiejar.New(rawURL)
	jar.SetCookies(u, cookies)
	client := http.Client{Jar: jar}
	resp, err := client.Get(rawURL)
    ...


## Link

- https://curl.haxx.se/docs/http-cookies.html

## Issue

- Can interconversion between `http.Cookie` and NetscapeCookieFormat ?

## TODO

- [ ] maintenance test
- [ ] complete main.go options

## License

MIT License
