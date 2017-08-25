package main

// Example:
// $ go run cmd/cookiedisk/main.go -v -b data/cookie.txt https://httpbin.org/cookies

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"

	"github.com/kosh04/cookiedisk"
)

type cli struct {
	cookies []*http.Cookie
	opt     struct {
		readFile  string
		writeFile string
		verbose   bool
	}
}

func init() {
	flag.Usage = func() {
		progname := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s [option] URL\n", progname)
		flag.PrintDefaults()
	}
}

func main() {
	var c cli
	flag.StringVar(&c.opt.readFile, "b", "", "`filename` to load cookies")
	flag.StringVar(&c.opt.writeFile, "c", "", "`filename` to save cookies")
	flag.BoolVar(&c.opt.verbose, "v", false, "verbose print")

	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if c.opt.readFile != "" {
		cookies, err := cookiedisk.ReadFile(c.opt.readFile)
		if err == nil {
			c.cookies = cookies
		}
	}

	urlStr := flag.Arg(0)
	res, err := c.send(urlStr)
	if err != nil {
		log.Fatal(err)
	}

	print(res)
}

func (c *cli) send(urlStr string) (string, error) {
	u, _ := url.Parse(urlStr)
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(u, c.cookies)

	client := http.Client{Jar: jar}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return "", err
	}
	if c.opt.verbose {
		for name, values := range req.Header {
			fmt.Printf("%s: %v\n", name, values)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("HTTP Status", resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	if c.opt.verbose {
		for name, values := range resp.Header {
			fmt.Printf("%s: %v\n", name, values)
		}
	}

	// update cookie-jar file
	c.cookies = jar.Cookies(u)
	if c.opt.writeFile != "" {
		err := cookiedisk.WriteFile(c.opt.writeFile, u, c.cookies)
		if err != nil {
			log.Print(err)
		}
	}

	return string(body), nil
}
