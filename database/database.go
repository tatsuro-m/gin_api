package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() {
	_, err := gorm.Open("postgres", "host=db user=postgres password=password")
	if err != nil {
		fmt.Println("connection failed!!!!!!!")
		fmt.Println("データベースの接続に失敗しました。。。")
		return
	}
	fmt.Println("DBへの接続には成功したっぽい")
}
