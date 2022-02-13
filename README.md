# ckecklist project
A small project that implements the checklist based on Gin and Grom frame

## front-end
The front-end part of the project is provided by Q1mi, see details：
https://github.com/Q1mi/go_web/blob/master/lesson25/dist.zip

## back-end
The back-end code is completed independently with reference to BiliBili video
BiliBili viedo:https://www.bilibili.com/video/BV1gJ411p7xC?p=27

## how to start
cd to main directory and excute `go run main.go`,then open the browser, enter 127.0.0.1:8000 in the address bar and press enter

## matters needing attention
In this project, gorm.model structure has been modified, and the ID field has added tag: json: "id", so that the keys of the front-end and back-end JSON can correspond one by one
~~~go
type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
~~~
