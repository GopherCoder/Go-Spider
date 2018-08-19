package cmd

import (
	"Go-Spider/infra/initial"
	"Go-Spider/src/model"
	"encoding/json"
	"fmt"
	"math/rand"

	"os"

	"time"

	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "get one image for database by choice",
	Run: func(cmd *cobra.Command, args []string) {
		initial.DBInit()
		defer initial.DataBase.Close()
		initial.DataBase.LogMode(false)
		switch args[1] {
		case "rand":

			var imageAddress model.ImageAddress
			var imageSize []model.ImageSize
			if dbError := initial.DataBase.Find(&imageSize).Error; dbError != nil {
				return
			}

			var sizeList []int
			for _, one := range imageSize {
				sizeList = append(sizeList, int(one.ID))
			}
			//rand.Seed(100)
			rand.Seed(time.Now().Unix())
			if dbError := initial.DataBase.Where("image_size_id = ?", sizeList[rand.Intn(len(sizeList))]).First(&imageAddress).Error; dbError != nil {
				fmt.Println("try again")
				return
			}
			jsonObject, _ := json.MarshalIndent(imageAddress, "", " ")
			fmt.Println(string(jsonObject))

		case "type":

			var imageAddress []model.ImageAddress
			if dbError := initial.DataBase.Where("size_type = ?", args[2]).Find(&imageAddress).Error; dbError != nil {
				fmt.Println("try again")
				return
			}
			rand.Seed(time.Now().Unix())
			number := rand.Intn(len(imageAddress))
			jsonObject, _ := json.MarshalIndent(imageAddress[number], "", " ")
			fmt.Println(string(jsonObject))
		case "download":
			// 随机选择图片，根据选择得到的图片进行下载
			// 下一步获取得到 N 张图片
		}
	},
}

func Execute() {
	if err := dbCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
