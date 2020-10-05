package database

import (
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	DisplayName string
	UserName    string
	TelegramID  int64 `gorm:"uniqueIndex:idx_telegram_id_bot_id"`
	BotID       int64 `gorm:"uniqueIndex:idx_telegram_id_bot_id"`
}

func (model *User) ToEntity(user *types.User) {
	user.BotID = model.BotID
	user.FirstName = model.FirstName
	user.LastName = model.LastName
	user.DisplayName = model.DisplayName
	user.TelegramID = model.TelegramID
	user.UserName = model.UserName
}

func (model *User) FromEntity(user *types.User) {
	model.BotID = user.BotID
	model.FirstName = user.FirstName
	model.LastName = user.LastName
	model.DisplayName = user.DisplayName
	model.TelegramID = user.TelegramID
	model.UserName = user.UserName
}
