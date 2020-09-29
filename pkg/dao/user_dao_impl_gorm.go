package dao

import (
	"errors"

	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserDaoImplGorm struct {
	db *gorm.DB
}

func NewUserDaoImplGorm(db *gorm.DB) UserDao {
	return &UserDaoImplGorm{
		db: db,
	}
}

func (dao *UserDaoImplGorm) Put(user *types.User) (*types.User, error) {

	resultingUser := &types.User{}
	userModel := &database.User{}
	copier.Copy(userModel, user)

	err := dao.db.Transaction(func(tx *gorm.DB) error {
		existingUser := &database.User{}
		if err := tx.Where(
			"bot_id=? AND telegram_id=?",
			user.BotID,
			user.TelegramID,
		).First(existingUser).Error; err != nil {
			//Error requesting existing user
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			//Else: user not found (is not actually an error)
			//We just need to create a user
			if err := tx.Create(userModel).Error; err != nil {
				return err
			}
			copier.Copy(resultingUser, userModel)
			return nil
		}
		//A user is found
		if err := tx.Model(existingUser).Updates(userModel).Error; err != nil {
			return err
		}
		copier.Copy(resultingUser, existingUser)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resultingUser, nil
}

func (dao *UserDaoImplGorm) Get(botID int64, telegramID int64) (*types.User, error) {
	userModel := &database.User{}

	result := dao.db.Where("bot_id=? AND telegram_id=?", botID, telegramID).First(userModel)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	user := &types.User{}
	copier.Copy(user, userModel)
	return user, nil

}
