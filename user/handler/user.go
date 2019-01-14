package handler

import (
	"cinema/user/db"
	pb "cinema/user/pb/user"
	"cinema/user/util"
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

type UserHandler struct {
	pb.User
}

func (u *UserHandler) CreateUser(ctx context.Context, user *pb.User, resp *pb.Response) error {
	hashPwd, err := util.GeneratePassword(user.Password, 1)
	if err != nil {
		return fmt.Errorf("generate password error: %v", err)
	}
	user.Password = hashPwd
	_, err = db.DB.Collection("user").InsertOne(context.Background(), &user)
	if err != nil {
		return err
	}
	resp.Success = true
	resp.Msg = "success"
	return nil
}

func (u *UserHandler) UserLogin(ctx context.Context, req *pb.LoginRequest, user *pb.LoginResponse) error {
	doc := bsonx.Doc{}
	db.DB.Collection().Clone()
	err := db.DB.Collection("user").FindOne(context.Background(),
		bson.D{{"username", req.Username}}).Decode(&doc)
	if err != nil {
		fmt.Println(err)
	}
	user.UserId = doc.Lookup("_id").ObjectID().Hex()
	user.Nickname = doc.Lookup("nickname").StringValue()
	user.Username = doc.Lookup("username").StringValue()
	user.Phone = doc.Lookup("phone").StringValue()
	i, _ := doc.Lookup("birthday").Int64OK()
	user.Birthday = i
	return nil
}
