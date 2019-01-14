package v1

import (
	"cinema/api/caech"
	apiClient "cinema/api/client"
	proto "cinema/api/proto/user"
	"cinema/api/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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
	_, err := us.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail",
			"code":    -1,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"code":    0,
	})
}

func UserLogin(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	resp, err := us.UserLogin(context.Background(), &proto.LoginRequest{
		Username: name,
		Password: password,
	})
	if err != nil {
		c.JSON(500, gin.H{"message": "fail"})
		return
	}
	token, err := util.GenerateToken(resp.Nickname)
	if err != nil {
		// 重试
		if token, err = util.GenerateToken(resp.Nickname); err != nil {
			panic("generate token error")
		}
	}
	// 设置缓存
	caech.Redis.Set(token, resp.UserId, time.Hour*24*3)
	caech.Redis.Del()
	fmt.Println(resp)
	c.JSON(200, gin.H{"token": token})

}
