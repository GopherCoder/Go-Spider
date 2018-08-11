package piaofang

import (
	"Go-Spider/infra/downloader"
	"Go-Spider/infra/errors"
	"fmt"
	"regexp"
)

// 使用正则表达式爬取的思路

type RankInfo struct {
	Rank            string `json:"rank"`
	MovieName       string `json:"movie_name"`
	StartYear       string `json:"start_year"`
	MovieType       string `json:"movie_type"`
	Director        string `json:"director"`
	BoxOfficeIncome string `json:"box_office_income"`
}

var (
	// <td class="num"><u>全球电影票房排行榜</u>第 1 名</td>
	rankPattern = `<tr><td class="num">(.*?)</td>`
	// <td class="title">《阿凡达》<span>Avatar</span></td>
	movieNamePattern = `<td class="title">(.*?)<span>(.*?)</span></td>`
	// <td class="year">2009</td>
	startYearPattern = `<td class="year">(.*?)</td>`
	// <td class="type">科幻/动作<span></span></td>
	movieTypePattern = `<td class="type">(.*?)<span></span></td>`
	// <td class="daoyan">詹姆斯·卡梅隆</td>
	directorPattern = `<td class="daoyan">(.*?)</td>`
	// <td class="piaofang"><span>2,787,965,087</span>$</td>
	boxOfficeIncomePattern = `<td class="piaofang"><span>(.*?)</span>(.*?)</td>`
)

func GetPFangRankInfo(url string) ([]RankInfo, error) {
	responseByte, err := downloader.GetHttpResponse(url, true)
	if err != nil {
		return nil, errors.ErrorPFangRankInfo
	}

	responseString := string([]byte(responseByte))
	var rank []string
	reRank := regexp.MustCompile(rankPattern)
	for index, subMatch := range reRank.FindAllStringSubmatch(responseString, -1) {
		if len(subMatch[1]) > 34 {
			subMatch[1] = subMatch[1][34:len(subMatch[1])]
		}
		fmt.Println(index, subMatch[1])
		rank = append(rank, subMatch[1])
	}

	var movieNameList []string
	movieName := regexp.MustCompile(movieNamePattern)
	for index, subMatch := range movieName.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, subMatch[1])
		movieNameList = append(movieNameList, subMatch[1])
	}

	var yearList []string
	startYear := regexp.MustCompile(startYearPattern)
	for index, subMatch := range startYear.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, subMatch[1])
		yearList = append(yearList, subMatch[1])
	}

	var movieTypeList []string
	movieType := regexp.MustCompile(movieTypePattern)
	for index, subMatch := range movieType.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, subMatch[1])
		movieTypeList = append(movieTypeList, subMatch[1])
	}

	var directorList []string
	director := regexp.MustCompile(directorPattern)
	for index, subMatch := range director.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, subMatch[1])
		directorList = append(directorList, subMatch[1])
	}

	var boxOfficeIncomeList []string
	boxOfficeIncome := regexp.MustCompile(boxOfficeIncomePattern)
	for index, subMatch := range boxOfficeIncome.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, subMatch[1]+"$")
		boxOfficeIncomeList = append(boxOfficeIncomeList, subMatch[1]+"$")
	}

	var result []RankInfo

	for index := 0; index < len(rank); index++ {
		var rankInfo RankInfo
		rankInfo = RankInfo{
			Rank:            rank[index],
			MovieName:       movieNameList[index],
			StartYear:       yearList[index],
			MovieType:       movieTypeList[index],
			Director:        directorList[index],
			BoxOfficeIncome: boxOfficeIncomeList[index],
		}
		result = append(result, rankInfo)

	}
	return result, nil

}
