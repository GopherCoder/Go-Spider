package githubtrending

import (
	"io/ioutil"
	"path/filepath"

	"bytes"

	"strings"

	"fmt"
	"regexp"

	"os"

	"encoding/json"

	"encoding/csv"

	"Go-Spider/infra/initial"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"github.com/parnurzeal/gorequest"
)

var Since = []string{"since", "weekly", "month"}

var Language = []string{"python", "go"}

var URL = "https://github.com/trending/%s?since=%s"

var FileBase, _ = filepath.Abs(filepath.Dir("./Go-Spider/domain/githubtrending/"))

var TypeList = []string{"text", "json", "csv"}

type Repositories struct {
	gorm.Model
	Name        string `gorm:"type:varchar" json:"name"`
	TotalStar   string `gorm:"type:varchar" json:"total_star"`
	SinceStar   string `gorm:"type:varchar" json:"since_star"`
	Fork        string `gorm:"type:varchar" json:"fork"`
	Description string `gorm:"type:varchar" json:"description"`
	SinceFlag   string `gorm:"type:varchar" json:"since_flag"`
	Language    string `gorm:"type:varchar" json:"language"`
}

type RepositoriesList []Repositories

var All RepositoriesList

func TrendingContent(language string, since string) RepositoriesList {
	url := fmt.Sprintf("https://github.com/trending/%s?since=%s", language, since)
	req := gorequest.New()
	response, _, _ := req.Get(url).End()
	defer response.Body.Close()

	responseBytes, _ := ioutil.ReadAll(response.Body)

	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(responseBytes))
	doc.Find("ol.repo-list li").Each(func(i int, selection *goquery.Selection) {
		name := strings.TrimSpace(selection.Find("div").Eq(0).Find("h3 a").Text())
		description := strings.TrimSpace(selection.Find("div").Eq(2).Find("p").Text())

		number := selection.Find("div").Eq(3)

		totalStar := strings.TrimSpace(number.Find("a").Eq(0).Text())

		fork := strings.TrimSpace(number.Find("a").Eq(1).Text())

		sinceStar := strings.TrimSpace(number.Find("span").Last().Text())

		pattern := "(.*?)stars"
		regexpPattern := regexp.MustCompile(pattern)
		all := regexpPattern.FindAllStringSubmatch(sinceStar, -1)
		var sinceNumber string
		for _, one := range all {
			//fmt.Println(sinceStar, one, one[1])
			sinceNumber = one[1]
		}

		//fmt.Println("name", name, "description", description, "totalStar", totalStar, "fork", fork, "sinceNumber", sinceNumber, "language", lanuage, "sinceFlag", since)
		data := Repositories{
			Name:        name,
			Description: description,
			TotalStar:   totalStar,
			Fork:        fork,
			SinceFlag:   since,
			Language:    language,
			SinceStar:   sinceNumber,
		}
		All = append(All, data)

	})
	return All

}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func SaveText(result RepositoriesList) {

	var file *os.File
	fileName := FileBase + "/trending.txt"

	if checkFileIsExist(fileName) {
		os.Remove(fileName)
	}

	file, _ = os.Create(fileName)
	for _, one := range result {
		string := fmt.Sprintf("Name: %s,  TotalStar: %s, SinceStar: %s, SinceFlag: %s, Language: %s, Description: %s",
			one.Name, one.TotalStar, one.SinceStar, one.SinceFlag, one.Language, one.Description)
		file.WriteString(string)
		file.WriteString("\n")
	}

	defer file.Close()
}

func SaveJson(result RepositoriesList) {

	var file *os.File
	fileName := FileBase + "/trending.json"

	if checkFileIsExist(fileName) {
		os.Remove(fileName)
	}

	file, _ = os.Create(fileName)
	byteResponse, _ := json.MarshalIndent(result, "", " ")
	file.WriteString(string(byteResponse))
	defer file.Close()

}

func SaveCsv(result RepositoriesList) {

	var f *os.File
	var err error
	fileName := FileBase + "/trending.csv"
	if checkFileIsExist(fileName) { //如果文件存在
		os.Remove(fileName)
	}
	f, err = os.Create(fileName) //创建文件
	if err != nil {
		f, err = os.Create(FileBase + "/trending.csv")
		fmt.Println(err, f.Name())
	}

	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流
	var header = []string{"name", "totalStar", "fork", "sinceNumber", "language", "since", "description"}
	w.Write(header)
	var data [][]string
	for _, one := range result {
		var singleString []string
		singleString = append(singleString, one.Name, one.TotalStar, one.Fork, one.SinceStar, one.Language, one.SinceFlag, one.Description)
		data = append(data, singleString)

	}

	w.WriteAll(data) //写入数据
	w.Flush()
}

func SavePostgres(result RepositoriesList) {
	for _, one := range result {
		if dbError := initial.DataBase.Save(&one).Error; dbError != nil {
			return
		}
	}
}

func SaveMarkDown(result RepositoriesList) {
	var file *os.File
	fileName := FileBase + "/trending.md"
	if checkFileIsExist(fileName) {
		os.Remove(fileName)
	}
	file, _ = os.Create(fileName)
	file.WriteString("## GithubTrending\n")
	for _, one := range result {
		title := fmt.Sprintf("### Name: %s\n", one.Name)
		file.WriteString(title)
		content := fmt.Sprintf("- Language: %s\n- Since: %s\n- SinceStar: %s\n- TotalSatr: %s\n- Description: %s\n", one.Language, one.SinceFlag, one.SinceStar, one.TotalStar, one.Description)
		file.WriteString(content)
	}
	defer file.Close()
}

func TrendingStart() {
	for _, one := range []string{"python", "go"} {
		for _, since := range []string{"day", "weekly", "month"} {
			TrendingContent(one, since)
		}
	}
	SaveText(All)
	SaveJson(All)
	SaveCsv(All)
	SavePostgres(All)
	SaveMarkDown(All)
}
