package routes

import (
	. "bubble/commons"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 主页
func IndexFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 添加待办
// 前端页面填写待办事项，然后发送POST请求到url，由该函数处理
func AddFunc(c *gin.Context) {
	// 从请求中提取数据
	var todo Todo
	c.BindJSON(&todo)

	// 将数据存入数据库
	err := DB.Create(&todo).Error

	// 向前端发送响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}

// 查看所有待办
func ViewAllFunc(c *gin.Context) {
	// 查询表中的所有数据
	var todoList []Todo
	err := DB.Find(&todoList).Error

	// 向前端发送数据
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// 查看某一个待办
func ViewOneFunc(c *gin.Context) {

}

// 修改某一个待办
func ChangeFunc(c *gin.Context) {
	// 从路径中获取id
	id, ok := c.Params.Get("id")
	if !ok {
		// 获取不到id
		c.JSON(http.StatusOK, gin.H{
			"error": "id not found",
		})
		return
	} else {
		// 获取到id，到数据库中查询
		var todo Todo
		err := DB.Where("id=?", id).First(&todo).Error
		if err != nil {
			// 查询失败
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
			return
		} else {
			// 查询成功
			// 开始更新
			c.BindJSON(&todo)
			err = DB.Save(&todo).Error
			if err != nil {
				// 更新失败
				c.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				// 更新成功
				c.JSON(http.StatusOK, gin.H{
					"message": "change ok",
				})
			}
		}
	}
}

// 删除某一个待办
func DeleteFunc(c *gin.Context) {
	// 从路径中获取id
	id, ok := c.Params.Get("id")
	if !ok {
		// 获取不到id
		c.JSON(http.StatusOK, gin.H{
			"error": "id not found",
		})
		return
	} else {
		// 获取到id，到数据库中删除
		err := DB.Where("id=?", id).Delete(Todo{}).Error
		if err != nil {
			// 删除失败
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
			return
		} else {
			// 删除成功
			c.JSON(http.StatusOK, gin.H{
				"message": "delete ok",
			})
		}
	}
}
