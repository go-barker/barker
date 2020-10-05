package database

import (
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type Bot struct {
	gorm.Model
	ID    int64 `gorm:"primaryKey"`
	Title string
	Token string
}

func (model *Bot) ToEntity(entity *types.Bot) {
	entity.ID = model.ID
	entity.Title = model.Title
	entity.Token = model.Token
}

func (model *Bot) FromEntity(entity *types.Bot) {
	model.ID = entity.ID
	model.Title = entity.Title
	model.Token = entity.Token
}
