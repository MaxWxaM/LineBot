package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper(fileName string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("./config")
	v.AddConfigPath("../config") //for debug
	v.SetConfigName(fileName)
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	return v, nil
}
