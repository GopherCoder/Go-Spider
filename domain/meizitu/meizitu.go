package meizitu

import (
	"fmt"

	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

/*
思路:
1. 获取所有图片的 URL 存储在 redis 内
2. 存储在 redis 内的类型是 list


*/

type Engine struct {
}

type Request struct {
	URL string
}

type Requests []Request

type Result struct {
	Item []interface{}
	Requests
}

func (e *Engine) Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		result := Worker(r)
		requests = append(requests, result.Requests...)
		for _, item := range result.Item {
			fmt.Println(item)
		}
	}

}

func getDoc(request Request) *goquery.Document {
	req := gorequest.New()
	response, _, _ := req.Get(request.URL).End()
	defer response.Body.Close()

	responseByte, _ := ioutil.ReadAll(response.Body)
	responseString := strings.NewReader(string([]byte(responseByte)))

	doc, _ := goquery.NewDocumentFromReader(responseString)
	return doc
}

func Worker(request Request) Result {
	doc := getDoc(request)
	requests := HomePage(doc)
	fmt.Println(requests)
	for _, one := range requests {
		ChildPage(one)
	}
	return Result{}

}

func HomePage(doc *goquery.Document) Requests {
	// 首页所有的子页的url
	var request Requests
	doc.Find("div.pic ul li").Each(func(i int, selection *goquery.Selection) {
		imageURL, _ := selection.Find("a").Attr("href")
		//fmt.Println(imageURL)
		request = append(request, Request{URL: imageURL})
	})
	return request

}

func ChildPage(request Request) {
	doc := getDoc(request)

	// #content > a > img
	childImage, _ := doc.Find("div#content a img").Attr("src")
	fmt.Println(childImage)

	// #page > a:nth-child(9)
	nextPage, ok := doc.Find("div#page a.ch.next").Attr("href")
	next := "http://www.mmjpg.com" + nextPage
	fmt.Println(nextPage)
	if ok {
		ChildPage(Request{URL: next})

	}

}

func Start() {
	// count 97
	var engine Engine
	engine.Run(Request{URL: "http://www.mmjpg.com/"})
	for index := 2; index <= 10; index++ {
		engine.Run(
			Request{
				URL: fmt.Sprintf("http://www.mmjpg.com/home/%d", index),
			},
		)
	}

}
