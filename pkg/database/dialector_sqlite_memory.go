package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDialectorSQLiteMemoryClient() gorm.Dialector {
	return sqlite.Open("file:client.db?mode=memory&cache=shared")
}

func NewDialectorSQLiteMemoryServer() gorm.Dialector {
	return sqlite.Open("file:server.db?mode=memory&cache=shared")
}
