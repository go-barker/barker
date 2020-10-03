package database

import (
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type Bot struct {
	gorm.Model
	types.Bot
	ID int64 `gorm:"primaryKey"`
}

type User struct {
	gorm.Model
	types.User
	TelegramID int64 `gorm:"uniqueIndex:idx_telegram_id_bot_id"`
	BotID      int64 `gorm:"uniqueIndex:idx_telegram_id_bot_id"`
}

type Campaign struct {
	gorm.Model
	types.Campaign
}

type Delivery struct {
	gorm.Model
	types.Delivery
	CampaignID int64               `gorm:"uniqueIndex:idx_campaign_bot_tg"`
	BotID      int64               `gorm:"uniqueIndex:idx_campaign_bot_tg"`
	TelegramID int64               `gorm:"uniqueIndex:idx_campaign_bot_tg"`
	State      types.DeliveryState `gorm:"index"`
}

func NewDatabase(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Campaign{})
	db.AutoMigrate(&Delivery{})
	db.AutoMigrate(&Bot{})
	return db.Debug(), nil
}
