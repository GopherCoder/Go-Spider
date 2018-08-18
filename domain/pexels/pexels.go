package pexels

import (
	"fmt"
	"regexp"

	"strconv"

	"Go-Spider/infra/initial"
	"Go-Spider/src/model"

	url2 "Go-Spider/infra/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

/*

思路:

1. 获取所有图片的链接,保存好size
2. 存储在数据库中
3. 根据size 随机从数据库中率选出 N 张进行下载
4. 存在 4 种样式：525*350(w*h，首页数据，高度一致，但长度不一), 1260*750(w*h), 940*650(w*h), 800*1200(w*h)
5. 设计两个表：一个图片大小数据表、一个图片地址表
*/

var valuePattern = `.*?&h=(.*?)&w=(.*?)$`

func getHeightWidth(value string) (int, int) {
	regexpHW := regexp.MustCompile(valuePattern)
	var height, width int
	for _, one := range regexpHW.FindAllStringSubmatch(value, -1) {
		fmt.Println(value, one[1], one[2])
		height, _ = strconv.Atoi(one[1])
		width, _ = strconv.Atoi(one[2])
	}
	return height, width

}

// 这个size 的图片地址是否存在
func imageSizeOperation(width, height string) (error, uint) {
	var imageSize model.ImageSize
	if dbError := initial.DataBase.Where("width = ? AND height = ?", width, height).First(&imageSize).Error; dbError != nil {
		imageSize.Width = width
		imageSize.Height = height
		if dbError := initial.DataBase.Save(&imageSize).Error; dbError != nil {
			return dbError, 0
		}
	}
	return nil, imageSize.ID
}

// 图片地址是否存在
func imageAddressOperation(address string, width string, height string, sizeType string) error {
	var imageAddress model.ImageAddress
	var imageSizeID uint
	var err error
	if dbError := initial.DataBase.Where("url = ?", address).First(&imageAddress).Error; dbError != nil {
		if err, imageSizeID = imageSizeOperation(width, height); err != nil {
			return err
		}
		imageAddress.URL = address
		imageAddress.ImageSizeID = imageSizeID
		imageAddress.SizeType = sizeType

		if dbError := initial.DataBase.Save(&imageAddress).Error; dbError != nil {
			return dbError
		}
	}
	return nil

}

func getPexelsImageLink(url string) {
	req := gorequest.New()
	resp, _, _ := req.Get(url).End()
	defer resp.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	//fmt.Println(doc.Html())
	doc.Find("div.photos article a.js-photo-link").Each(func(i int, selection *goquery.Selection) {
		/*
			<img srcset="https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&amp;cs=tinysrgb&amp;h=350 1x, https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&amp;cs=tinysrgb&amp;dpr=2&amp;h=350 2x" width="525" height="350" style="background:rgb(212,213,213)" class="photo-item__img" alt="Free stock photo of building, architecture, lines, apartment" data-big-src="https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&amp;cs=tinysrgb&amp;h=750&amp;w=1260" data-large-src="https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&amp;cs=tinysrgb&amp;h=650&amp;w=940" data-pin-media="https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&amp;cs=tinysrgb&amp;fit=crop&amp;h=1200&amp;w=800" src="https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&amp;cs=tinysrgb&amp;h=350">
		*/

		// 1. 图片大小数据表
		// 2. 图片地址数据表

		width, _ := selection.Find("img").Attr("width")
		height, _ := selection.Find("img").Attr("height")
		//fmt.Println(fmt.Sprintf("width: %s, height: %s", width, height))

		if err, _ := imageSizeOperation(width, height); err != nil {
			return
		}

		defaultImage, _ := selection.Find("img").Attr("src")
		//fmt.Println(defaultImage)
		if err := imageAddressOperation(defaultImage, width, height, "default"); err != nil {
			return
		}

		bigImage, _ := selection.Find("img").Attr("data-big-src")
		bigHeight, bigWidth := getHeightWidth(bigImage)
		//fmt.Println(bigImage, bigHeight, bigWidth)

		if err := imageAddressOperation(bigImage, strconv.Itoa(bigHeight), strconv.Itoa(bigWidth), "big"); err != nil {
			return
		}

		largeImage, _ := selection.Find("img").Attr("data-large-src")
		largeHeight, largeWidth := getHeightWidth(largeImage)
		//fmt.Println(largeImage, largeHeight, largeWidth)

		if err := imageAddressOperation(largeImage, strconv.Itoa(largeHeight), strconv.Itoa(largeWidth), "large"); err != nil {
			return
		}

		mediaImage, _ := selection.Find("img").Attr("data-pin-media")
		mediaHeight, mediaWidth := getHeightWidth(mediaImage)
		//fmt.Println(mediaImage, mediaHeight, mediaWidth)

		if err := imageAddressOperation(mediaImage, strconv.Itoa(mediaHeight), strconv.Itoa(mediaWidth), "media"); err != nil {
			return
		}

	})

}

func Start() {
	url := url2.PexlesURL
	for index := 1; index < 10; index++ {
		newURL := fmt.Sprintf(url, index)
		getPexelsImageLink(newURL)
	}

}
