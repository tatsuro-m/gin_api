package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connect() (db *gorm.DB, err error) {
	dsn := "host=" + os.Getenv("DB_HOST") + " port=5432" + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"
	return gorm.Open("postgres", dsn)
}

func Init() {
	_, err := connect()

	if err != nil {
		fmt.Println(" DB connection failed!!!!!!!")

		for i := 0; i < 3; i++ {
			fmt.Println("retry:" + strconv.Itoa(i+1))
			time.Sleep(time.Second * 5)
			_, err := connect()
			if err == nil {
				break
			}
		}

		return
	}
	fmt.Println("DB connection success!!!!")
}
