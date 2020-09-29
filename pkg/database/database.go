package database

import (
	"github.com/corporateanon/barker/pkg/config"
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/driver/mysql"
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
	ID int64 `gorm:"primaryKey"`
}

type Campaign struct {
	gorm.Model
	types.Campaign
}

type Delivery struct {
	gorm.Model
	types.Delivery
	CampaignID int64               `gorm:"uniqueIndex:idx_campaign_user"`
	UserID     int64               `gorm:"uniqueIndex:idx_campaign_user"`
	State      types.DeliveryState `gorm:"index"`
}

func NewDatabase(config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DBConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Campaign{})
	db.AutoMigrate(&Delivery{})
	db.AutoMigrate(&Bot{})
	return db.Debug(), nil
}
