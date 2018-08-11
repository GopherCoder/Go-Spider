package gushiwen

import (
	"Go-Spider/infra/downloader"
	"strings"

	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func getDoc(url string) (*goquery.Document, error) {
	response, err := downloader.GetResponseByRequest(url, "GET", "", "")
	if err != nil {
		return nil, err
	}

	responseString := strings.NewReader(string([]byte(response)))
	return goquery.NewDocumentFromReader(responseString)

}

// 正则表达式获取作品数目
func getNumber(value string) uint {
	var number uint
	numberPattern := "► (.*?)篇.*?"
	numberCompile := regexp.MustCompile(numberPattern)
	for _, one := range numberCompile.FindAllStringSubmatch(value, -1) {
		if len(one) < 1 {
			return 0
		}
		numberTemp, _ := strconv.Atoi(one[1])
		number = uint(numberTemp)
		break

	}
	return uint(number)

}
