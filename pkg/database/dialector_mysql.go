package database

import (
	"github.com/corporateanon/barker/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDialectorMySQL(config *config.Config) gorm.Dialector {
	return mysql.Open(config.DBConnection)
}
