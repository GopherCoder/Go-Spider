package downloader

import (
	"Go-Spider/infra/errors"
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
)

func GetResponseByRequest(url string, method string, body string, userAgent string) ([]byte, error) {

	request := gorequest.New()

	var (
		resp gorequest.Response
		err  []error
		//requestBody string
	)

	if method == "GET" {
		resp, _, err = request.Get(url).
			Set("Content-Type", "application/json").
			Set("User-Agent", userAgent).
			End()
	} else {
		resp, _, err = request.Post(url).
			Type("multipart").
			Send(body).
			End()
	}
	if len(err) != 0 {
		return nil, errors.ErrorResponse
	}
	defer resp.Body.Close()
	response, error := ioutil.ReadAll(resp.Body)
	return response, error
}
