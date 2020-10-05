package database

import (
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	ID      int64
	BotID   int64 `gorm:"index"`
	Title   string
	Message string
	Active  bool
}

func (model *Campaign) ToEntity(entity *types.Campaign) {
	entity.ID = model.ID
	entity.Active = model.Active
	entity.BotID = model.BotID
	entity.Message = model.Message
	entity.Title = model.Title
}

func (model *Campaign) FromEntity(entity *types.Campaign) {
	model.ID = entity.ID
	model.Active = entity.Active
	model.BotID = entity.BotID
	model.Message = entity.Message
	model.Title = entity.Title
}
