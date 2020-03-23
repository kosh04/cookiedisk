package cookiedisk_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/kosh04/cookiedisk"
)

func ExampleReadFile() {
	rawURL := "http://httpbin.org/cookies"
	cookies, _ := cookiedisk.ReadFile("testdata/cookie.txt")
	u, _ := url.Parse(rawURL)
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(u, cookies)
	client := http.Client{Jar: jar}
	resp, err := client.Get(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	// Output:
	// {
	//   "cookies": {
	//     "k1": "v1"
	//   }
	// }
}
