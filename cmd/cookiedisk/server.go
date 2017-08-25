package main

import (
	"fmt"
	"net/http"
	"time"
)

func cookiesDump(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Cookies())
}

func cookiesSetFunc(cookies []*http.Cookie) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, c := range cookies {
			http.SetCookie(w, c)
		}
	}
}

func cookiesDeleteFunc(name string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    name,
			Expires: time.Time{},
		})
	}
}
