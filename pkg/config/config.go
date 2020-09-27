package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver     string `validate:"required,oneof=mysql sqlite"`
	DBConnection string `validate:"required"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	v := viper.New()
	v.AutomaticEnv()
	c.DBDriver = v.GetString("db_driver")
	c.DBConnection = v.GetString("db_connection")
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
