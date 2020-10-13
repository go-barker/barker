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

func (this *DeliveryDaoImplGorm) Take(botID int64, campaignID int64) (*dao.DeliveryTakeResult, error) {
	errNoRecepients := errors.New("no recepients")

	result := &dao.DeliveryTakeResult{
		Delivery: &types.Delivery{},
		User:     &types.User{},
		Campaign: &types.Campaign{},
	}

	err := this.db.Transaction(func(tx *gorm.DB) error {
		type resultWrapper struct {
			database.User
			CampaignID int64
		}

		resultModel := &resultWrapper{}

		query := tx.
			Table("users").
			Select("users.*", "campaigns.id as campaign_id").
			Joins("inner join campaigns on "+
				"campaigns.bot_id = users.bot_id "+
				"AND (campaigns.id = ? OR 0 = ?)", campaignID, campaignID).
			Joins(
				"left outer join deliveries on "+
					"deliveries.telegram_id = users.telegram_id "+
					"AND deliveries.bot_id = users.bot_id "+
					"AND deliveries.campaign_id = campaigns.id",
			).
			// Joins(
			// 	"left outer join deliveries on "+
			// 		"deliveries.telegram_id = users.telegram_id "+
			// 		"AND deliveries.bot_id = users.bot_id "+
			// 		"AND deliveries.campaign_id = ?",
			// 	campaignID).
			Where("deliveries.telegram_id IS NULL").
			Where("users.deleted_at IS NULL").
			Where("users.bot_id = ?", botID).
			Order("campaigns.created_at DESC").
			Limit(1).
			Scan(resultModel)

		if err := query.Error; err != nil {
			return err
		}
		if resultModel.ID == 0 {
			return errNoRecepients
		}

		deliveryModel := &database.Delivery{
			CampaignID: resultModel.CampaignID,
			BotID:      botID,
			TelegramID: resultModel.TelegramID,
			State:      types.DeliveryStateProgress,
		}
		if err := tx.Create(deliveryModel).Error; err != nil {
			return err
		}

		campaignModel := &database.Campaign{}
		if err := tx.Where("id = ?", resultModel.CampaignID).Find(campaignModel).Error; err != nil {
			return err
		}
		campaignModel.ToEntity(result.Campaign)
		deliveryModel.ToEntity(result.Delivery)
		resultModel.ToEntity(result.User)

		return nil
	})

	if err != nil {
		if errors.Is(err, errNoRecepients) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
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
