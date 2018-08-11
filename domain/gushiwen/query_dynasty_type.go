package gushiwen

import (
	"Go-Spider/infra/initial"
	"Go-Spider/src/model"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

//  朝代
func parseDynasty() {
	common(true)

}

// 诗的类型
func parsePoetryType() {
	common(false)
}

func common(ok bool) {
	url := "https://www.gushiwen.org/shiwen/"
	doc, _ := getDoc(url)
	if ok {
		doc.Find("div.right div.sons").Eq(2).Find("div.cont a").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
			var dynasty model.Dynasty
			dynasty.Name = selection.Text()
			initial.DataBase.Save(&dynasty)
		})
	} else {
		doc.Find("div.right div.sons").Eq(3).Find("div.cont a").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
			var poetType model.PoetryType
			poetType.TypeName = selection.Text()
			initial.DataBase.Save(&poetType)
		})

	}

}
