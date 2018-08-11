package maoyan

import (
	"Go-Spider/infra/downloader"
	"Go-Spider/infra/errors"
	"fmt"
	"regexp"
	"strings"
)

type RankListMaoYanInfo struct {
	Rank              string `json:"rank"`
	MovieName         string `json:"movie_name"`
	StartYear         string `json:"start_year"`
	BoxOfficeIncome   string `json:"box_office_income"`
	AvgEachTicket     string `json:"avg_each_ticket"`
	PerformancePeople string `json:"performance_people"`
	Url               string `json:"url"`
}

var base64ToMap = make(map[string]string)

func init() {
	base64ToMap = map[string]string{
		"&#xf0f2": "0",
		"&#xe9bd": "1",
		"&#xeb2f": "2",
		"&#xe5ce": "3",
		"&#xf58a": "4",
		"&#xf77a": "5",
		"&#xe6e3": "6",
		"&#xe0b0": "7",
		"&#xeb7a": "8",
		"&#xe8af": "9",
	}

}

var (
	maoYanRootURl                = "https://piaofang.maoyan.com/"
	rankPattern                  = `<li class="col0">(.*?)</li>`
	movieNameAndStartYearPattern = `<li class="col1">
            <p class="first-line">(.*?)</p>
            <p class="second-line">(.*?)上映</p>
        </li>`
	boxOfficeIncomePattern   = `<li class="col2 tr"><i class="cs">(.*?)</i></li>`
	avgEachTicketPattern     = `<li class="col3 tr"><i class="cs">(.*?)</i></li>`
	performancePeoplePattern = `<li class="col4 tr"><i class="cs">(.*?)</i></li>`
	urlPattern               = `<ul class="row" data-com="hrefTo,href:'(.*?)'">`
)

func GetMaoYanRankInfo(url string) ([]RankListMaoYanInfo, error) {
	response, err := downloader.GetHttpResponse(url, false)
	if err != nil {
		return nil, errors.ErrorMaoYanRankList
	}
	responseString := string([]byte(response))
	fmt.Println(responseString)

	var urlList []string
	urlRegexp := regexp.MustCompile(urlPattern)
	for index, one := range urlRegexp.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, one)
		urlList = append(urlList, maoYanRootURl+one[1])
	}

	var rankList []string
	rankRegexp := regexp.MustCompile(rankPattern)

	for index, rank := range rankRegexp.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, rank)
		if index > 0 {
			rankList = append(rankList, rank[1])
		}

	}

	var movieNameList []string
	var startYearList []string
	movieNameAndStartYearRegexp := regexp.MustCompile(movieNameAndStartYearPattern)

	for index, movie := range movieNameAndStartYearRegexp.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, movie)
		movieNameList = append(movieNameList, movie[1])
		startYearList = append(startYearList, movie[2])
	}

	var boxOfficeIncomeList []string
	boxOfficeIncomeRegexp := regexp.MustCompile(boxOfficeIncomePattern)

	for index, boxOffice := range boxOfficeIncomeRegexp.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, boxOffice)
		boxList := strings.Split(boxOffice[1], ";")
		for _, box := range boxList {
			var value string
			value += base64ToMap[box]
			boxOfficeIncomeList = append(boxOfficeIncomeList, value)
		}

	}

	var avgEachTicketList []string
	avgEachTicketRegexp := regexp.MustCompile(avgEachTicketPattern)

	for index, avgEach := range avgEachTicketRegexp.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, avgEach)
		avgEachTicketList = append(avgEachTicketList, avgEach...)
	}

	var performancePeoPleList []string
	performancePeopleRegexp := regexp.MustCompile(performancePeoplePattern)

	for index, performance := range performancePeopleRegexp.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, performance)
		performancePeoPleList = append(performancePeoPleList, performance...)
	}

	var result []RankListMaoYanInfo
	for index := 0; index < len(movieNameList); index++ {
		var OneRankListMaoYanInfo RankListMaoYanInfo
		OneRankListMaoYanInfo = RankListMaoYanInfo{
			Rank:      rankList[index],
			MovieName: movieNameList[index],
			StartYear: startYearList[index],
			Url:       urlList[index],
		}
		result = append(result, OneRankListMaoYanInfo)

	}

	return result, nil
}
