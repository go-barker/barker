package database

import (
	"gorm.io/gorm"
)

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
