package controllers

import (
	"basic/task4/database"
	"basic/task4/filters"
	"basic/task4/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Post(ginServer *gin.Engine) {
	//连接数据库
	conn := database.DBConnection()
	// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容
	ginServer.POST("/createPost", filters.JWTAuth(), func(context *gin.Context) {
		// GetRawData	原理，就是从c.Request.Body中获取数据，并将数据返回 []byte
		data, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		userID := m["userId"]
		title := m["title"]
		content := m["content"]
		// 插入数据库 注册成功
		user := models.Post{Title: title.(string), Content: content.(string), UserID: uint(userID.(float64))}
		conn.Create(&user)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "create success",
			"data": m,
		})
	})
	//实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
	ginServer.POST("/getPost", filters.JWTAuth(), func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		userID := m["userId"]
		fmt.Println(userID)
		if userID != nil {
			var results []map[string]interface{}
			conn.Table("posts").Where("user_id = ? and deleted_at is null", userID).Find(&results)
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "query success",
				"data": results,
			})
		} else {
			var results []map[string]interface{}
			conn.Table("posts").Where("deleted_at is null").Find(&results)
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "query success",
				"data": results,
			})
		}
	})

	//实现文章的更新功能，只有文章的作者才能更新自己的文章。
	ginServer.POST("/updatePost", filters.JWTAuth(), func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		var posts models.Post
		_ = json.Unmarshal(data, &m)
		id := m["id"]
		user_id := m["user_id"]
		m["updated_at"] = time.Now()
		conn.Debug().Model(&posts).Where("id = ? and user_id = ?", id, user_id).Updates(m)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "update success",
			"data": m,
		})

	})

	//实现文章的删除功能，只有文章的作者才能删除自己的文章。
	ginServer.POST("/deletePost", filters.JWTAuth(), func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		var posts models.Post
		_ = json.Unmarshal(data, &m)
		id := m["id"]
		user_id := m["user_id"]
		m["updated_at"] = time.Now()
		conn.Debug().Model(&posts).Where("id = ? and user_id = ?", id, user_id).Delete(&posts)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "delete success",
			"data": m,
		})
	})

}
