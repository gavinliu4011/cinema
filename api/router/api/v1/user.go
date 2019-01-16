package v1

import (
	"cinema/api/caech"
	apiClient "cinema/api/client"
	proto "cinema/api/proto/user"
	"cinema/api/util"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	us = apiClient.UserClient()
)

func RegistryUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	user := &proto.User{
		Username: username,
		Password: password,
		Nickname: nickname,
	}
	resp, err := us.CreateUser(context.Background(), user)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "error",
			"code":    -1,
		})
		return
	}
	result := gin.H{}
	if resp.Success {
		result["message"] = "success"
		result["code"] = 0
	} else {
		result["message"] = resp.Msg
		result["code"] = -1
	}
	c.JSON(200, result)
}

func UserLogin(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	resp, err := us.UserLogin(context.Background(), &proto.LoginRequest{
		Username: name,
		Password: password,
	})
	if err != nil {
		c.JSON(500, gin.H{"message": "fail"})
		return
	}
	expiration := time.Hour * 24 * 3
	token, err := util.GenerateToken(resp.Nickname, expiration)
	// 设置缓存
	caech.Redis.Set(token, resp.UserId, expiration)
	c.JSON(200, gin.H{"token": token})

}
