package module

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// 放映厅
type ProjectionRoom struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name"`       // 放映厅名字
	Ticket    uint8              `json:"ticket"`     // 总票数
	Status    uint8              `json:"status"`     // 状态 0:空闲 1:放映中
	MovieName string             `json:"movie_name"` // 播放电影名
	StartTime time.Time          `json:"start_time"`
	EndTime   time.Time          `json:"end_time"`
}

// 初始化放映厅
func InitRoom() {
	_ = DB.Collection("rooms").Drop(context.Background())
	var rooms []interface{}
	for i := 1; i <= 5; i++ {
		p := ProjectionRoom{
			Name:      fmt.Sprintf("放映厅-%v", i),
			Ticket:    200,
			Status:    0,
			ID:        primitive.NewObjectID(),
			StartTime: time.Now(),
		}
		rooms = append(rooms, p)
	}
	result, err := DB.Collection("rooms").InsertMany(context.Background(), rooms)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result)
	cursor, err := DB.Collection("rooms").Find(context.Background(), bson.D{})
	doc := bsonx.Doc{}
	for cursor.Next(context.Background()) {
		doc = doc[:0]
		err := cursor.Decode(&doc)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(doc)
	}
}
