package controllers

import (
	"basic/task4/database"
	"basic/task4/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Comment(ginServer *gin.Engine) {
	//连接数据库
	conn := database.DBConnection()
	//实现评论的创建功能，已认证的用户可以对文章发表评论。
	ginServer.POST("/createComment", func(context *gin.Context) {
		// GetRawData    原理，就是从c.Request.Body中获取数据，并将数据返回 []byte
		data, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		userID := m["user_id"]
		postID := m["post_id"]
		content := m["content"]
		//fmt.Println(userID)
		//fmt.Println(postID)
		//fmt.Println(content)
		// 插入数据库
		comment := models.Comment{Content: content.(string), UserID: uint(userID.(float64)), PostID: uint(postID.(float64))}
		conn.Create(&comment)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "create success",
			"data": m,
		})

	})
	// 实现评论的读取功能，支持获取某篇文章的所有评论列表
	ginServer.POST("/getComments", func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		postID := m["post_id"]
		// 接受数据结构体
		type results struct {
			UserID    uint
			PostID    uint
			Title     string
			Content   string
			CommentId uint
			Comments  string
		}
		var result []results
		conn.Model(&models.Post{}).Select("posts.user_id,posts.id as post_id,posts.title,posts.content,comments.id as comments_id,comments.content as comments ").
			Joins("left join comments on  posts.id = comments.post_id").Where("posts.id = ?", postID).Scan(&result)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "query success",
			"data": result,
		})
	})
	//查询全部
	ginServer.GET("/getAllComments", func(context *gin.Context) {
		var comments []models.Comment
		conn.Find(&comments)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "query success",
			"data": comments,
		})

	})

}
