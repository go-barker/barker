package dao

import (
	"errors"

	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type DeliveryDaoImplGorm struct {
	db          *gorm.DB
	campaignDao CampaignDao
	userDao     UserDao
}

func NewDeliveryDaoImplGorm(
	db *gorm.DB,
	campaignDao CampaignDao,
	userDao UserDao,
) DeliveryDao {
	return &DeliveryDaoImplGorm{
		db:          db,
		campaignDao: campaignDao,
		userDao:     userDao,
	}
}

func (dao *DeliveryDaoImplGorm) Take(botID int64, campaignID int64) (*types.Delivery, *types.User, error) {
	errNoRecepients := errors.New("no recepients")

	resultingDelivery := &types.Delivery{}
	resultingUser := &types.User{}

	err := dao.db.Transaction(func(tx *gorm.DB) error {
		recepientUser := &database.User{}
		query := tx.
			Table("users").
			Select("users.*").
			Joins(
				"left outer join deliveries on "+
					"deliveries.telegram_id = users.telegram_id "+
					"AND deliveries.bot_id = users.bot_id "+
					"AND deliveries.campaign_id = ?",
				campaignID,
			).
			Where("deliveries.telegram_id IS NULL").
			Where("users.deleted_at IS NULL").
			Where("users.bot_id = ?", botID).
			Limit(1).
			Scan(recepientUser)

		if err := query.Error; err != nil {
			return err
		}
		if recepientUser.ID == 0 {
			return errNoRecepients
		}

		deliveryModel := &database.Delivery{
			CampaignID: campaignID,
			BotID:      botID,
			TelegramID: recepientUser.TelegramID,
			State:      types.DeliveryStateProgress,
		}
		if err := tx.Create(deliveryModel).Error; err != nil {
			return err
		}

		copier.Copy(resultingDelivery, deliveryModel)
		copier.Copy(resultingUser, recepientUser)

		return nil
	})

	if err != nil {
		if errors.Is(err, errNoRecepients) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return resultingDelivery, resultingUser, nil
}

func (dao *DeliveryDaoImplGorm) Success(campaignID int64, userID int64) error {
	panic("not implemented")
}

func (dao *DeliveryDaoImplGorm) Fail(campaignID int64, userID int64) error {
	panic("not implemented")
}
