package initial

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var DataBase *gorm.DB

func DBInit() {
	db, err := gorm.Open("postgres", GetPostGreConfig())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("database login success")
	db.LogMode(true)
	DataBase = db
}

func GetPostGreConfig() string {
	viper.SetConfigName("settings")
	viper.AddConfigPath("$GOPATH/src/Go-Spider/configs")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var (
		dbName   string
		port     string
		user     string
		sslMode  string
		password string
		host     string
	)

	dbName = viper.GetString("db.dbname")
	port = viper.GetString("db.port")
	user = viper.GetString("db.user")
	sslMode = viper.GetString("db.sslmode")
	password = viper.GetString("db.password")
	host = viper.GetString("db.host")
	fmt.Println(fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s", host, user, dbName, port, sslMode, password))
	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s", host, user, dbName, port, sslMode, password)

}
