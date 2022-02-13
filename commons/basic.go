package commons

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var DB *gorm.DB

func InitDB() (err error) {
	addr := "root:@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", addr)
	if err != nil {
		return
	}
	err = DB.DB().Ping()
	return
}
