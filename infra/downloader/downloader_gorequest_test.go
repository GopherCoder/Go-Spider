package downloader

import (
	"Go-Spider/infra/url"
	"fmt"
	"testing"
)

func TestGetResponseByRequest(t *testing.T) {

	var xinBangBody string

	xinBangBody = `{"end":"2018-08-09","rank_name":"%E6%97%B6%E4%BA%8B","rank_name_group":"%E8%B5%84%E8%AE%AF","start":"2018-08-09","nonce":"6b6a45639","xyz":"ed5f089de3e9229f4884a78c2de0c73f"}`

	tests := []struct {
		url       string
		method    string
		body      string
		userAgent string
	}{
		{
			url:       url.XinBangURL,
			method:    "POST",
			body:      xinBangBody,
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		},
		//}, {
		//	url:    url.GuShiWenURL,
		//	method: "GET",
		//	body:   "",
		//},
	}

	for _, test := range tests {
		response, err := GetResponseByRequest(test.url, test.method, test.body, test.userAgent)
		fmt.Println(string([]byte(response)), err)
	}
}
