package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// テストすべきはファイルに保存したcookieをサーバに渡せるかどうか

func TestCookies(t *testing.T) {
	cookies := []*http.Cookie{{Name: "k1", Value: "v1"}}
	cli := cli{cookies: cookies}

	ts := httptest.NewServer(http.HandlerFunc(cookiesDump))
	defer ts.Close()

	_, err := cli.send(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cli.cookies, cookies) {
		t.Errorf("cookie=%v, want=%v", cli.cookies, cookies)
	}
}

func TestSetCooies(t *testing.T) {
	cookies := []*http.Cookie{
		{Name: "k1", Value: "v1"},
		{Name: "k2", Value: "v2"},
	}
	handler := cookiesSetFunc(cookies)

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	cli := cli{}
	_, err := cli.send(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cli.cookies, cookies) {
		t.Errorf("cookie=%v, want=%v", cli.cookies, cookies)
	}
}
