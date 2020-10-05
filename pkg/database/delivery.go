package database

import (
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	CampaignID int64               `gorm:"uniqueIndex:idx_campaign_bot_tg"`
	BotID      int64               `gorm:"uniqueIndex:idx_campaign_bot_tg"`
	TelegramID int64               `gorm:"uniqueIndex:idx_campaign_bot_tg"`
	State      types.DeliveryState `gorm:"index"`
}

func (model *Delivery) ToEntity(entity *types.Delivery) {
	entity.BotID = model.BotID
	entity.CampaignID = model.CampaignID
	entity.State = model.State
	entity.TelegramID = model.TelegramID
}

func (model *Delivery) FromEntity(entity *types.Delivery) {
	model.BotID = entity.BotID
	model.CampaignID = entity.CampaignID
	model.State = entity.State
	model.TelegramID = entity.TelegramID
}
