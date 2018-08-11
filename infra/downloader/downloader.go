package downloader

import (
	"Go-Spider/infra/errors"
	"fmt"
	"io/ioutil"
	"net/http"

	URL "Go-Spider/infra/url"

	"bufio"
	"io"

	"strings"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func GetHttpResponse(url string, ok bool) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorRequest
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

	client := http.DefaultClient

	response, err := client.Do(request)

	if err != nil {
		return nil, errors.ErrorResponse
	}

	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.ErrorStatusCode
	}
	if ok {

		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	} else {
		return ioutil.ReadAll(response.Body)
	}

}

func determineCharset(i io.Reader) encoding.Encoding {
	resp, err := bufio.NewReader(i).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(resp, "")
	fmt.Println(e)
	return e
}

func GetHttpResponseBySelenium(url string) (string, error) {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		},
	}
	caps.AddChrome(chromeCaps)

	service, err := selenium.NewChromeDriverService(
		URL.ChromeDriver, 9515,
	)

	defer service.Stop()

	if err != nil {
		return "", errors.ErrorSelenium
	}

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))

	if err != nil {
		return "", errors.ErrorWebDriver
	}

	err = webDriver.Get(url)
	return webDriver.PageSource()

}

// PostHttpResponse ...
func PostHttpResponse(url string, body string, ok bool) ([]byte, error) {

	payload := strings.NewReader(body)
	requests, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, errors.ErrorRequest
	}
	requests.Header.Add("X-Requested-With", "XMLHttpRequest")
	requests.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	requests.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(requests)
	if err != nil {
		return nil, errors.ErrorResponse
	}

	fmt.Println(response.StatusCode)

	defer response.Body.Close()
	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}
	return ioutil.ReadAll(response.Body)

}
