package config

import (
	"context"
	"fmt"
	"time"

	"github.com/MaxWxaM/linebot/internal/app/model/client"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientConfig struct {
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	UserName       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	DialTimeoutSec int    `mapstructure:"dial_timeout_sec"`
	MaxOpenConns   int    `mapstructure:"max_open_conns"` // if this is 0, it will be set to math.MaxInt64
}

type LintBotClientConfig struct {
	Secret string `mapstructure:"secret"`
	Token  string `mapstructure:"token"`
	UserId string `mapstructure:"userId"`
}

type HttpConfig struct {
	Port int
}

func NewMongoClient(cfg ClientConfig) (*mongo.Client, func(), error) {
	credential := options.Credential{
		Username: cfg.UserName,
		Password: cfg.Password,
	}
	option := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port))
	option.SetAuth(credential)
	option.SetConnectTimeout(time.Duration(cfg.DialTimeoutSec) * time.Second)
	option.SetMaxPoolSize(uint64(cfg.MaxOpenConns))
	client, err := mongo.Connect(context.TODO(), option)

	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}

	return client, cleanup, nil
}

func NewMongoClientConfig() (ClientConfig, error) {
	var (
		err error
		o   = new(ClientConfig)
	)

	v, _ := NewViper("app")

	if err = v.UnmarshalKey("mongo_client", o); err != nil {
		return ClientConfig{}, err
	}

	return *o, err
}

func NewHttpOptions() *HttpConfig {
	var (
		err error
		o   = new(HttpConfig)
	)

	v, _ := NewViper("app")

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil
	}

	return o
}

func NewLineBotClientConfig() (LintBotClientConfig, error) {
	var (
		err error
		o   = new(LintBotClientConfig)
	)

	v, _ := NewViper("app")

	if err = v.UnmarshalKey("linebot_client", o); err != nil {
		return LintBotClientConfig{}, err
	}

	return *o, err
}

func NewLineBotClient(cfg LintBotClientConfig) (*client.CustomizeLineBotClient, error) {
	bot, err := linebot.New(cfg.Secret, cfg.Token)

	if err != nil {
		return nil, err
	}

	return &client.CustomizeLineBotClient{
		Client: bot,
		UserId: cfg.UserId,
	}, nil
}
