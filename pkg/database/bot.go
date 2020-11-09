package database

import (
	"time"

	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type Bot struct {
	gorm.Model
	ID              int64 `gorm:"primaryKey"`
	Title           string
	Token           string    `gorm:"uniqueIndex"`
	RRAccessTime    time.Time `gorm:"index"`
	RRPossiblyEmpty bool      `gorm:"index"`
}

func (model *Bot) ToEntity(entity *types.Bot) {
	entity.ID = model.ID
	entity.Title = model.Title
	entity.Token = model.Token
	entity.RRAccessTime = model.RRAccessTime
	entity.RRPossiblyEmpty = model.RRPossiblyEmpty
}

func (model *Bot) FromEntity(entity *types.Bot) {
	model.ID = entity.ID
	model.Title = entity.Title
	model.Token = entity.Token
	model.RRAccessTime = entity.RRAccessTime
	model.RRPossiblyEmpty = entity.RRPossiblyEmpty
}
