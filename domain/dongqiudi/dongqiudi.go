package dongqiudi

import (
	"Go-Spider/infra/errors"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

/*

1. 懂球帝 APP 数据抓取分析实例
2. FIFA 排行: 国家队、俱乐部、球员
3. 返回格式 json

*/

var (
	FiFaRankingURL = "https://sport-data.dongqiudi.com/soccer/biz/data/fifa_ranking?type=nation&&app=dqd&version=162&platform=android"
	ClubRankingURL = "https://sport-data.dongqiudi.com/soccer/biz/data/fifa_ranking?type=club&&app=dqd&version=162&platform=android"
	MarketValueURL = "https://sport-data.dongqiudi.com/soccer/biz/data/market_value_ranking?&app=dqd&version=162&platform=android"
	UrlList        = []string{FiFaRankingURL, ClubRankingURL, MarketValueURL}
)

func dongqiudiGetDoc(url string) (*goquery.Document, error) {
	req := gorequest.New()

	resp, _, err := req.Get(url).
		Set("Content-Type", "application/json").
		End()

	if err != nil {
		return nil, errors.ErrorResponse
	}

	return goquery.NewDocumentFromReader(resp.Body)

}

type DongQiuDiEngine struct {
	URL           string
	ParseFunction func(document *goquery.Document)
}

func (d *DongQiuDiEngine) Run() {
	doc, _ := dongqiudiGetDoc(d.URL)
	d.ParseFunction(doc)

}

func getJsonValue(document *goquery.Document) gjson.Result {
	result := gjson.Parse(document.Text())
	return result
}

func common(result gjson.Result) {
	headers := result.Get("content.header").Array()
	for _, one := range headers {
		fmt.Println(one.String())
	}
	data := result.Get("content.data").Array()
	for _, one := range data {
		currentData := one.Get("current_data").String()
		lastRank := one.Get("last_rank").String()
		previousData := one.Get("previous_data").String()
		rank := one.Get("rank").String()
		teamID := one.Get("team_id").String()
		teamLogo := one.Get("team_logo").String()
		teamName := one.Get("team_name").String()
		fmt.Println(currentData, lastRank, previousData, rank, teamID, teamLogo, teamName)
	}
}

func ParseFiFaRanking(document *goquery.Document) {
	result := getJsonValue(document)
	common(result)
}

func ParseClubRanking(document *goquery.Document) {
	result := getJsonValue(document)
	common(result)

}

func ParseMarketValue(document *goquery.Document) {
	result := getJsonValue(document)
	headers := result.Get("content.header").Array()
	for _, one := range headers {
		fmt.Println(one.String())
	}
	data := result.Get("content.data").Array()
	for _, one := range data {
		count := one.Get("count").String()
		currency := one.Get("currency").String()
		personID := one.Get("person_id").String()
		personLogo := one.Get("person_logo").String()
		personName := one.Get("person_name").String()
		rank := one.Get("rank").String()
		teamID := one.Get("team_id").String()
		teamLogo := one.Get("team_logo").String()
		teamName := one.Get("team_name").String()
		value := one.Get("value").String()
		fmt.Println(count, currency, personID, personLogo, personName, rank, teamID, teamLogo, teamName, value)
	}

}

func StartDongQiuDi() {
	var DongQiuDiList []DongQiuDiEngine
	var funcList = []func(document *goquery.Document){ParseFiFaRanking, ParseClubRanking, ParseMarketValue}
	for index, One := range UrlList {
		DongQiuDiList = append(DongQiuDiList, DongQiuDiEngine{
			URL:           One,
			ParseFunction: funcList[index],
		})

	}

	for _, One := range DongQiuDiList {
		One.Run()
	}
}
