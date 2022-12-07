package repo

import (
	"context"
	"time"

	"github.com/MaxWxaM/linebot/internal/app/model/input"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMsgRepository interface {
	LogMsg(ctx context.Context, msg input.Msg) error
}

type MongoMsgRepo struct {
	db *mongo.Database
}

func NewMongoMsgRepo(cli *mongo.Client) IMsgRepository {
	db := cli.Database("Msg")

	return &MongoMsgRepo{
		db: db,
	}
}

func (repo MongoMsgRepo) LogMsg(ctx context.Context, msg input.Msg) error {
	collection := repo.db.Collection("MsgLog")
	_, err := collection.InsertOne(ctx, concatMsgLog(msg))
	return err
}

func concatMsgLog(msg input.Msg) bson.D {
	return bson.D{
		{Key: "T", Value: primitive.Timestamp{T: uint32(time.Now().Unix())}},
		{Key: "Header", Value: msg.Header},
		{Key: "Content", Value: msg.Content},
	}
}
