package main

import (
	"Go-Spider/domain/dongqiudi"
	"Go-Spider/domain/gushiwen"
	"Go-Spider/infra/initial"
	"Go-Spider/src/model"
	"fmt"
)

func init() {
	//initial.DBInit()
	//CreateTable()
	//Start()
}

func CreateTable() {
	initial.DataBase.AutoMigrate(
		&model.RankListCBOInfo{},
		&model.Dynasty{},
		&model.PoetryType{},
		&model.Poet{},
		&model.PoetryInfo{},
	)
}
func Start() {
	{
		//cbooo.SaveCBOIntoDB()
	}

	{
		// 古诗文抓取
		var engine gushiwen.SimpleEngine
		var results []gushiwen.Request
		var rootURL = "https://www.gushiwen.org/shiwen/default.aspx?page=%d&type=0&id=0"
		for index := 1; index <= 50; index++ {
			url := fmt.Sprintf(rootURL, index)
			results = append(results, gushiwen.Request{
				URL: url,
			})
		}
		engine.Run(
			results,
		)

	}

}

func main() {
	//defer initial.DataBase.Close()
	//meizitu.Start()
	dongqiudi.StartDongQiuDi()

}
