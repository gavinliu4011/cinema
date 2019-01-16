package handler

import (
	"cinema/user/module"
	pb "cinema/user/pb/user"
	"cinema/user/util"
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

type UserHandler struct {
}

// 创建一个用户
func (u *UserHandler) CreateUser(ctx context.Context, user *pb.User, resp *pb.Response) error {
	hashPwd, err := util.GeneratePassword(user.Password, 1)
	if err != nil {
		return fmt.Errorf("generate password error: %v", err)
	}
	coll := module.DB.Collection("user")
	doc := bsonx.Doc{}
	_ = coll.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&doc)
	if len(doc) != 0 {
		resp.Success = false
		resp.Msg = "用户已存在"
		return nil
	}

	user.Password = hashPwd
	_, err = coll.InsertOne(context.Background(), &user)
	module.DB.Collection("")
	if err != nil {
		return err
	}
	resp.Success = true
	resp.Msg = "success"
	return nil
}

// 用户登陆
func (u *UserHandler) UserLogin(ctx context.Context, req *pb.LoginRequest, user *pb.LoginResponse) error {
	doc := bsonx.Doc{}
	err := module.DB.Collection("user").FindOne(context.Background(),
		bson.D{{"username", req.Username}}).Decode(&doc)
	if err != nil {
		return err
	}
	validate := util.ValidatePassword(doc.Lookup("password").String(), req.Password)
	if !validate {
		return fmt.Errorf("密码错误")
	}
	user.UserId = doc.Lookup("_id").ObjectID().Hex()
	user.Nickname = doc.Lookup("nickname").StringValue()
	user.Username = doc.Lookup("username").StringValue()
	user.Phone = doc.Lookup("phone").StringValue()
	i, _ := doc.Lookup("birthday").Int64OK()
	user.Birthday = i
	return nil
}
