package gushiwen

import (
	"fmt"

	"Go-Spider/src/model"

	"strconv"

	"Go-Spider/infra/initial"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Request struct {
	URL string
}

type SimpleEngine struct {
}

func (s *SimpleEngine) Run(seeds []Request) {

	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		Worker(r.URL, parsePoetry)
	}
}

func Worker(url string, f func(document *goquery.Document)) {
	doc, err := getDoc(url)
	if err != nil {
		return
	}
	f(doc)

}

func parsePoetry(doc *goquery.Document) {

	doc.Find("div.left div.sons").Each(func(i int, selection *goquery.Selection) {
		onePart := selection.Find("div.cont")
		title := onePart.Find("p a b").Text()
		dynasty := onePart.Find("p.source a").Eq(0).Text()
		author := onePart.Find("p.source  a").Eq(1).Text()
		authorURL, _ := onePart.Find("p.source a").Eq(1).Attr("href")
		content := strings.TrimSpace(onePart.Find("div.contson").Text())
		twoPart := selection.Find("div.tool")
		liked, _ := strconv.Atoi(strings.TrimSpace(twoPart.Find("div.good a span").Text()))
		//fmt.Println(title, dynasty, author, content, authorURL, liked)

		docAuthor, _ := getDoc(authorURL)
		var poet model.Poet
		poet = parsePoet(docAuthor)

		var dynastyModel model.Dynasty
		if dbError := initial.DataBase.Where("name = ?", dynasty).Find(&dynastyModel).Error; dbError != nil {
			dynastyModel.Name = dynasty
			initial.DataBase.Save(&dynastyModel)
		}

		poet.DynastyID = dynastyModel.ID

		fmt.Println("poet", poet, author)
		//诗人
		var findPoet model.Poet
		if dbError := initial.DataBase.Where("name = ?", author).First(&findPoet).Error; dbError != nil {
			initial.DataBase.Save(&poet)
		}

		var poetry model.PoetryInfo
		poetry = model.PoetryInfo{
			Title:     title,
			DynastyID: dynastyModel.ID,
			Liked:     uint(liked),
			Content:   content,
			PoetID:    poet.ID,
		}
		fmt.Println("poetry", poetry, "like", liked)

		if dbError := initial.DataBase.Save(&poetry).Error; dbError != nil {
			return
		}

	})

}

// 诗人
func parsePoet(doc *goquery.Document) model.Poet {
	var poet model.Poet

	author := doc.Find("div.titleleft").Parent().Find("b").Text()
	if !doc.Find("div").HasClass("sonspic") {
		return model.Poet{
			Name: author,
		}
	}
	doc.Find("div.sonspic").Each(func(i int, selection *goquery.Selection) {
		onePart := selection.Find("div.cont")
		image, _ := onePart.Find("div.divimg a img").Attr("src")
		content := onePart.Find("p").Eq(1).Text()
		numberTemp := onePart.Find("p").Eq(1).Find("a").Text()
		number := getNumber(numberTemp)
		twoPart := selection.Find("div.tool")
		liked, _ := strconv.Atoi(strings.TrimSpace(twoPart.Find("div.good a span").Text()))
		//fmt.Println(image, content, number, liked)

		poet = model.Poet{
			ImageURL:    image,
			Description: content,
			Number:      uint(number),
			Liked:       uint(liked),
			Name:        author,
		}

	})
	return poet
}
