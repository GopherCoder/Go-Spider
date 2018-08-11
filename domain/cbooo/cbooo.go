package cbooo

import (
	"Go-Spider/infra/downloader"
	"Go-Spider/infra/initial"
	"Go-Spider/infra/url"
	"fmt"

	"strings"

	"Go-Spider/src/model"

	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

var areaMap = make(map[string]string)
var urlList []string
var areaList []string
var total int64

func init() {
	areaMap = map[string]string{
		"中国":   "50",
		"中国香港": "37",
		"中国台湾": "40",
		"美国":   "1",
		"英国":   "25",
		"德国":   "16",
		"法国":   "4",
		"日本":   "30",
		"加拿大":  "2",
		"意大利":  "7",
	}

	for key, value := range areaMap {
		var oneUrl string
		oneUrl = fmt.Sprintf(url.CboRealURL, value, 1)
		urlList = append(urlList, oneUrl)
		areaList = append(areaList, key)
	}

}

func SaveCBOIntoDB() {
	getCboRankList()
}

type RankListCBOInfo struct {
	BoxOffice    string `json:"box_office"`
	MovieID      string `json:"movie_id"`
	MovieEnName  string `json:"movie_en_name"`
	MovieName    string `json:"movie_name"`
	Ranking      string `json:"ranking"`
	DefaultImage string `json:"default_image"`
	ReleaseYear  string `json:"release_year"`
	Country      string `json:"country"`
}

func getCboRankList() {

	for indexArea, oneUrl := range urlList {
		time.Sleep(1 * time.Millisecond * 1000 * 10)
		response, err := downloader.GetHttpResponse(oneUrl, false)
		if err != nil {
			return
		}

		responseString := string([]byte(response))
		totalPage := gjson.Parse(responseString).Get("tPage").Int()
		total += gjson.Parse(responseString).Get("tCount").Int()
		fmt.Println("totalPage", totalPage, "current_count", gjson.Parse(responseString).Get("tCount").Int(), "total:", total)
		fmt.Println(oneUrl)

		for index := 1; index <= int(totalPage); index++ {
			newOneUrl := strings.Replace(oneUrl, "pIndex=1", fmt.Sprintf("pIndex=%d", index), -1)
			fmt.Println(newOneUrl)
			go func(url string) {
				time.Sleep(1 * time.Millisecond * 1000 * 10)
				newResponse, err := downloader.GetHttpResponse(newOneUrl, false)
				if err != nil {

					//return
				}
				//fmt.Println(string([]byte(newResponse)))
				newResponseString := string([]byte(newResponse))
				info := gjson.Parse(newResponseString).Get("pData")

				for _, oneInfo := range info.Array() {
					var oneRank model.RankListCBOInfo
					oneRank = model.RankListCBOInfo{
						Ranking:      oneInfo.Get("Ranking").String(),
						MovieID:      oneInfo.Get("ID").String(),
						MovieEnName:  oneInfo.Get("MovieEnName").String(),
						MovieName:    oneInfo.Get("MovieName").String(),
						ReleaseYear:  oneInfo.Get("releaseYear").String(),
						DefaultImage: oneInfo.Get("defaultImage").String(),
						BoxOffice:    oneInfo.Get("BoxOffice").String(),
						Country:      areaList[indexArea],
					}
					fmt.Println(oneRank)
					initial.DataBase.Create(&oneRank)
				}
			}(newOneUrl)
		}
	}
	fmt.Println("total", total)

}

func getAreaList(url string) {
	response, err := downloader.GetHttpResponse(url, false)
	if err != nil {
		return
	}
	stringReader := strings.NewReader(string([]byte(response)))
	document, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		return
	}

	document.Find("#selArea option").Each(func(i int, selection *goquery.Selection) {
		country := selection.Text()
		value, _ := selection.Attr("value")
		fmt.Println(country, value)

	})

	document.Find("#selType option").Each(func(i int, selection *goquery.Selection) {
		if i > 0 {
			MovieType := selection.Text()
			MovieTypeNumber, _ := selection.Attr("value")
			fmt.Println(MovieType, MovieTypeNumber)

		}

	})
}

type RealTimeMovie struct {
	Rank         string `json:"rank"`
	MovieName    string `json:"movie_name"`
	BoxOffice    string `json:"box_office"`
	Rate         string `json:"rate"`
	SumBoxOffice string `json:"sum_box_office"`
	TicketRate   string `json:"ticket_rate"`
	Day          string `json:"day"`
}

func getRealTimeList(url string) {
	response, err := downloader.GetHttpResponse(url, false)
	if err != nil {
		return
	}

	responseString := string([]byte(response))
	//fmt.Println(responseString)

	document, err := goquery.NewDocumentFromReader(strings.NewReader(responseString))
	if err != nil {
		return
	}

	document.Find("#topdatatr tr").Each(func(i int, selection *goquery.Selection) {
		var RealTimeTicket RealTimeMovie
		RealTimeTicket = RealTimeMovie{
			Rank:         selection.Find("td").Eq(0).Text(),
			MovieName:    selection.Find("td").Eq(1).Text(),
			BoxOffice:    selection.Find("td").Eq(2).Text(),
			Rate:         selection.Find("td").Eq(3).Text(),
			SumBoxOffice: selection.Find("td").Eq(4).Text(),
			TicketRate:   selection.Find("td").Eq(5).Text(),
			Day:          selection.Find("td").Eq(6).Text(),
		}

		fmt.Println(RealTimeTicket)

	})

}
