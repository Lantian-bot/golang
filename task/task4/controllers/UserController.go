package controllers

import (
	"basic/task4/database"
	"basic/task4/filters"
	"basic/task4/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func User(ginServer *gin.Engine) {
	//连接数据库
	conn := database.DBConnection()
	// 用户注册
	ginServer.POST("/registerUser", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		email := context.PostForm("email")
		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		password = string(hashedPassword)
		// 验证用户名是否存在 如果用户名或邮箱已存在，则返回冲突状态码
		var user models.User
		conn.Where("username = ?", username).First(&user)
		result := conn.Where("email = ?", email).First(&user)
		if user.ID != 0 || result.RowsAffected == 1 {
			context.JSON(http.StatusConflict, gin.H{"error": "Username or Email already exists"})
		} else {
			// 插入数据库 注册成功
			user := models.User{Username: username, Password: password, Email: email}
			conn.Create(&user)
			context.JSON(http.StatusOK, gin.H{
				"code":     http.StatusOK,
				"msg":      "registere success",
				"username": username,
				"password": password,
				"email":    email,
			})
		}

	})
	// 用户登录
	ginServer.POST("/loginUser", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		// 验证用户名是否存在
		var user models.User
		conn.Where("username = ?", username).First(&user)
		if user.ID == 0 {
			context.JSON(http.StatusNotFound, gin.H{"error": "Username not found"})
			return
		}
		// 验证密码是否正确
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}
		// 生成 JWT
		token, err := filters.GenerateToken(strconv.Itoa(int(user.ID)), []string{username})
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		}
		context.JSON(http.StatusOK, gin.H{
			"code":     http.StatusOK,
			"msg":      "login success",
			"username": username,
			"password": user.Password,
			"token":    token,
		})

	})
	// 查询用户列表
	ginServer.GET("/userList", filters.JWTAuth(), func(context *gin.Context) {
		var users []models.User
		conn.Find(&users)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "query success",
			"data": users,
		})
	})

}
