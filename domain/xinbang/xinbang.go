package xinbang

import (
	"Go-Spider/infra/downloader"
	"fmt"
)

func GetXinBangReport(url string) {

	body := "end=2018-08-07&rank_name=%E6%97%B6%E4%BA%8B&rank_name_group=%E8%B5%84%E8%AE%AF&start=2018-08-07&nonce=9aff0f999&xyz=4a24bc7445929ba9da5afdeae8fab3be"

	response, err := downloader.PostHttpResponse(url, body, false)
	if err != nil {
		return
	}

	fmt.Println(string([]byte(response)))
}

func SaveWithText() {}

func SaveWithCsv() {}

func SaveWithDataBase() {}
