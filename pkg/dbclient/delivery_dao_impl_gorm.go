package dbclient

import (
	"errors"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type DeliveryDaoImplGorm struct {
	db          *gorm.DB
	campaignDao dao.CampaignDao
	userDao     dao.UserDao
}

func NewDeliveryDaoImplGorm(
	db *gorm.DB,
	campaignDao dao.CampaignDao,
	userDao dao.UserDao,
) dao.DeliveryDao {
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
		recepientUserModel := &database.User{}
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
			Joins("inner join campaigns on "+
				"campaigns.bot_id = users.bot_id "+
				"AND campaigns.id = ?", campaignID).
			Where("deliveries.telegram_id IS NULL").
			Where("users.deleted_at IS NULL").
			Where("users.bot_id = ?", botID).
			Limit(1).
			Scan(recepientUserModel)

		if err := query.Error; err != nil {
			return err
		}
		if recepientUserModel.ID == 0 {
			return errNoRecepients
		}

		deliveryModel := &database.Delivery{
			CampaignID: campaignID,
			BotID:      botID,
			TelegramID: recepientUserModel.TelegramID,
			State:      types.DeliveryStateProgress,
		}
		if err := tx.Create(deliveryModel).Error; err != nil {
			return err
		}

		deliveryModel.ToEntity(resultingDelivery)
		recepientUserModel.ToEntity(resultingUser)

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

func (dao *DeliveryDaoImplGorm) SetState(delivery *types.Delivery, state types.DeliveryState) error {
	if state != types.DeliveryStateProgress &&
		state != types.DeliveryStateSuccess &&
		state != types.DeliveryStateFail {
		return errors.New("Wrong delivery state")
	}

	return dao.db.Model(&database.Delivery{}).
		Where("bot_id = ? AND campaign_id = ? AND telegram_id = ?",
			delivery.BotID,
			delivery.CampaignID,
			delivery.TelegramID).
		Update("state", state).Error
}

func (dao *DeliveryDaoImplGorm) GetState(delivery *types.Delivery) (types.DeliveryState, error) {
	result := &struct{ State types.DeliveryState }{}
	query := dao.db.Model(&database.Delivery{}).
		Where("bot_id = ? AND campaign_id = ? AND telegram_id = ?",
			delivery.BotID,
			delivery.CampaignID,
			delivery.TelegramID).
		Select("state").Scan(result)
	if query.Error != nil {
		return 0, query.Error
	}
	return result.State, nil
}
