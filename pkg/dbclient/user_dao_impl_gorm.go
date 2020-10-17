package dbclient

import (
	"errors"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/pagination"
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type UserDaoImplGorm struct {
	db *gorm.DB
}

func NewUserDaoImplGorm(db *gorm.DB) dao.UserDao {
	return &UserDaoImplGorm{
		db: db,
	}
}

func (dao *UserDaoImplGorm) Put(user *types.User) (*types.User, error) {

	resultingUser := &types.User{}
	userModel := &database.User{}
	userModel.FromEntity(user)

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
			userModel.ToEntity(resultingUser)
			return nil
		}
		//A user is found
		if err := tx.Model(existingUser).Updates(userModel).Error; err != nil {
			return err
		}
		existingUser.ToEntity(resultingUser)
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
	userModel.ToEntity(user)
	return user, nil
}

func (dao *UserDaoImplGorm) List(botID int64, pageRequest *types.PaginatorRequest) ([]types.User, *types.PaginatorResponse, error) {
	userModelsList := []database.User{}
	db := dao.db.Table("users").
		Order("created_at DESC").
		Where("bot_id = ?", botID)
	resp := pagination.Paging(&pagination.Param{
		DB:    db,
		Page:  int(pageRequest.Page),
		Limit: int(pageRequest.Size),
	}, &userModelsList)

	if err := db.Error; err != nil {
		return nil, nil, err
	}

	usersList := make([]types.User, len(userModelsList))
	for i, model := range userModelsList {
		model.ToEntity(&usersList[i])
	}
	return usersList,
		&types.PaginatorResponse{
			Page:  resp.Page,
			Size:  resp.Limit,
			Total: resp.TotalPage,
		},
		nil
}
