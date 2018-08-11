package nuomi

import (
	"Go-Spider/infra/downloader"
	"fmt"

	"github.com/tidwall/gjson"
)

type CommonRankListFiled struct {
	MovieId   string `json:"movie_id"`
	MovieName string `json:"movie_name"`
	Index     string `json:"index"`
}

type BoxRankListNuoMiInfo struct {
	CommonRankListFiled
	RealTimeBox string `json:"real_time_box"`
}

type FocusRankListNuoMiInfo struct {
	CommonRankListFiled
	MovieType string `json:"movie_type"`
	FocusNum  string `json:"focus_num"`
}

type SearchRankListNuoMiInfo struct {
	CommonRankListFiled
	MovieType   string `json:"movie_type"`
	SearchCount string `json:"search_count"`
}

type ReputationRankList struct {
	CommonRankListFiled
	MovieType  string `json:"movie_type"`
	Reputation string `json:"reputation"`
}

type AgeRankList struct {
	CommonRankListFiled
	MovieType string `json:"movie_type"`
	Score     string `json:"score"`
}

type AnnualBoxMovieRankList struct {
	CommonRankListFiled
	MovieType string `json:"movie_type"`
}

type AnnualReputationMovieRankList struct {
	AgeRankList
}

type RegionRankList struct {
	AgeRankList
}

func GetRankNuoMiInfo(url string) {
	response, err := downloader.GetHttpResponse(url, false)
	if err != nil {
		return
	}

	//fmt.Println(string([]byte(response)))

	box := gjson.Parse(string([]byte(response))).Get("boxRankList")

	var boxes []BoxRankListNuoMiInfo
	for _, OneBoxRank := range box.Array() {
		var box BoxRankListNuoMiInfo
		box = BoxRankListNuoMiInfo{
			RealTimeBox: OneBoxRank.Get("realTimeBox").String(),
			CommonRankListFiled: CommonRankListFiled{
				MovieId:   OneBoxRank.Get("movieId").String(),
				MovieName: OneBoxRank.Get("movieName").String(),
				Index:     OneBoxRank.Get("index").String(),
			},
		}
		boxes = append(boxes, box)
	}

	focus := gjson.Parse(string([]byte(response))).Get("focusRankList")

	var focusRankLists []FocusRankListNuoMiInfo
	for _, OneFocus := range focus.Array() {
		var OneFocusRankList FocusRankListNuoMiInfo
		OneFocusRankList = FocusRankListNuoMiInfo{
			CommonRankListFiled: CommonRankListFiled{
				MovieId:   OneFocus.Get("movieId").String(),
				MovieName: OneFocus.Get("movieName").String(),
				Index:     OneFocus.Get("index").String(),
			},
			MovieType: OneFocus.Get("movieType").String(),
			FocusNum:  OneFocus.Get("focusNum").String(),
		}
		focusRankLists = append(focusRankLists, OneFocusRankList)
	}

	search := gjson.Parse(string([]byte(response))).Get("searchRankList")
	var searchRankLists []SearchRankListNuoMiInfo
	for _, OneSearch := range search.Array() {
		var oneSearchRankList SearchRankListNuoMiInfo
		oneSearchRankList = SearchRankListNuoMiInfo{
			CommonRankListFiled: CommonRankListFiled{
				MovieId:   OneSearch.Get("movieId").String(),
				MovieName: OneSearch.Get("movieName").String(),
				Index:     OneSearch.Get("index").String(),
			},
			SearchCount: OneSearch.Get("searchCount").String(),
		}
		searchRankLists = append(searchRankLists, oneSearchRankList)

	}

	reputationRankList := gjson.Parse(string([]byte(response))).Get("reputationRankList")

	var reputationRankLists []ReputationRankList
	for _, oneReputation := range reputationRankList.Array() {
		var oneReputationRankList ReputationRankList
		oneReputationRankList = ReputationRankList{
			CommonRankListFiled: CommonRankListFiled{
				MovieId:   oneReputation.Get("movieId").String(),
				MovieName: oneReputation.Get("movieName").String(),
				Index:     oneReputation.Get("index").String(),
			},
			Reputation: oneReputation.Get("reputation").String(),
		}
		reputationRankLists = append(reputationRankLists, oneReputationRankList)
	}

	ageRankList := gjson.Parse(string([]byte(response))).Get("ageRankList")
	var ageRankLists []AgeRankList
	for _, OneAgeRank := range ageRankList.Array() {
		var oneAgeRankList AgeRankList
		oneAgeRankList = AgeRankList{
			CommonRankListFiled: CommonRankListFiled{
				MovieId:   OneAgeRank.Get("movieId").String(),
				MovieName: OneAgeRank.Get("movieName").String(),
				Index:     OneAgeRank.Get("index").String(),
			},
			Score: OneAgeRank.Get("score").String(),
		}
		ageRankLists = append(ageRankLists, oneAgeRankList)
	}

	annualBoxMovie := gjson.Parse(string([]byte(response))).Get("annualBoxMovieRankList")

	var annualBoxMovies []AnnualBoxMovieRankList
	for _, OneAnnualBox := range annualBoxMovie.Array() {
		var OneAnnualRank AnnualBoxMovieRankList
		OneAnnualRank = AnnualBoxMovieRankList{
			CommonRankListFiled: CommonRankListFiled{
				MovieId:   OneAnnualBox.Get("movieId").String(),
				MovieName: OneAnnualBox.Get("movieName").String(),
				Index:     OneAnnualBox.Get("index").String(),
			},
			MovieType: OneAnnualBox.Get("movieType").String(),
		}
		annualBoxMovies = append(annualBoxMovies, OneAnnualRank)

	}
	annualReputation := gjson.Parse(string([]byte(response))).Get("annualReputationMovieRankList")
	var annualReputations []AnnualReputationMovieRankList
	for _, OneAnnual := range annualReputation.Array() {
		var OneAnnualRank AnnualReputationMovieRankList
		OneAnnualRank = AnnualReputationMovieRankList{
			AgeRankList: AgeRankList{
				CommonRankListFiled: CommonRankListFiled{
					MovieId:   OneAnnual.Get("movieId").String(),
					MovieName: OneAnnual.Get("movieName").String(),
					Index:     OneAnnual.Get("index").String(),
				},
				Score:     OneAnnual.Get("score").String(),
				MovieType: OneAnnual.Get("movieType").String(),
			},
		}
		annualReputations = append(annualReputations, OneAnnualRank)
	}
	regionRank := gjson.Parse(string([]byte(response))).Get("regionRankList")

	var regionRanks []RegionRankList
	for _, OneRegion := range regionRank.Array() {
		var OneRank RegionRankList
		OneRank = RegionRankList{
			AgeRankList: AgeRankList{
				CommonRankListFiled: CommonRankListFiled{
					MovieId:   OneRegion.Get("movieId").String(),
					MovieName: OneRegion.Get("movieName").String(),
					Index:     OneRegion.Get("index").String(),
				},
				Score:     OneRegion.Get("score").String(),
				MovieType: OneRegion.Get("movieType").String(),
			},
		}
		regionRanks = append(regionRanks, OneRank)

	}

	fmt.Println(boxes, focusRankLists, searchRankLists, reputationRankLists, ageRankLists, annualBoxMovies,
		annualReputations, regionRanks)

}
