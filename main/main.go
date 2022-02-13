package main

import (
	. "bubble/commons"
	"bubble/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 连接数据库
	err := InitDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	// 自动迁移
	DB.AutoMigrate(&Todo{})

	// 创建服务器
	r := gin.Default()

	// 静态文件路径，第一个参数是文件路径，第二个参数是HTML里面写的路径
	// r.Static("../static", "static")
	r.StaticFS("/static", http.Dir("../static"))
	r.LoadHTMLFiles("../templates/index.html") // 解析HTML

	// 主页
	r.GET("/", routes.IndexFunc)

	// v1 待办事项
	v1Group := r.Group("/v1")
	{
		//// 添加
		// 添加待办事项
		v1Group.POST("/todo", routes.AddFunc)

		//// 查看
		// 查看所有待办
		v1Group.GET("/todo", routes.ViewAllFunc)
		// 查看某一个待办
		v1Group.GET("/todo/:id", routes.ViewOneFunc)

		//// 修改
		// 修改某一个待办
		v1Group.PUT("/todo/:id", routes.ChangeFunc)

		// 删除
		// 删除某一个待办
		v1Group.DELETE("/todo/:id", routes.DeleteFunc)
	}
	r.Run("127.0.0.1:8000")
}
