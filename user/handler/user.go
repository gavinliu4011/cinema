package handler

import (
	"cinema/user/module"
	pb "cinema/user/pb/user"
	"cinema/user/util"
	"context"
	"encoding/json"
	"fmt"

	"github.com/getsentry/raven-go"
	"github.com/micro/go-micro/broker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	userCreateTopic = "com.cinema.srv.user.created"
	userLoginTopic  = "com.cinema.srv.user.login"
)

type UserHandler struct {
}

// 创建一个用户
func (u *UserHandler) CreateUser(ctx context.Context, user *pb.User, resp *pb.Response) error {
	hashPwd, err := util.GeneratePassword(user.Password, 1)
	if err != nil {
		raven.CaptureError(err, nil)
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
	result, err := coll.InsertOne(context.Background(), &user)
	module.DB.Collection("")
	if err != nil {
		return err
	}
	resp.Success = true
	resp.Msg = "success"

	id, _ := result.InsertedID.(primitive.ObjectID)
	b, _ := json.Marshal(user)
	msg := broker.Message{
		Header: map[string]string{"userId": id.Hex()},
		Body:   []byte(b),
	}
	err = broker.Publish(userCreateTopic, &msg)
	if err != nil {
		raven.CaptureError(err, nil)
	}
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
	userId := doc.Lookup("_id").ObjectID().Hex()
	user.UserId = userId
	user.Nickname = doc.Lookup("nickname").StringValue()
	user.Username = doc.Lookup("username").StringValue()
	user.Phone = doc.Lookup("phone").StringValue()
	i, _ := doc.Lookup("birthday").Int64OK()
	user.Birthday = i
	b, _ := json.Marshal(&user)
	msg := broker.Message{
		Header: map[string]string{"userId": userId},
		Body:   []byte(b),
	}
	err = broker.Publish(userLoginTopic, &msg)
	if err != nil {
		raven.CaptureError(err, nil)
	}
	return nil
}
