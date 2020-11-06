package dbclient

import (
	"errors"
	"time"

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

func (this *DeliveryDaoImplGorm) Take(botID int64, campaignID int64, telegramID int64) (*dao.DeliveryTakeResult, error) {
	result := &dao.DeliveryTakeResult{
		Delivery: &types.Delivery{},
		User:     &types.User{},
		Campaign: &types.Campaign{},
	}

	recipientsNotFound := false

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
			Where("deliveries.telegram_id IS NULL").
			Where("users.deleted_at IS NULL").
			Where("campaigns.deleted_at IS NULL").
			Where("campaigns.active = true").
			Where("users.bot_id = ?", botID).
			Order("campaigns.created_at DESC").
			Limit(1)
		if telegramID != 0 {
			query = query.Where("users.telegram_id = ?", telegramID)
		}
		query = query.Scan(resultModel)

		if err := query.Error; err != nil {
			return err
		}
		if resultModel.ID == 0 {
			recipientsNotFound = true
			if err := this.updateBotPossiblyEmptyStatus(tx, botID, true); err != nil {
				return err
			}
			return nil
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

		if err := this.updateBotPossiblyEmptyStatus(tx, botID, false); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	if recipientsNotFound {
		return nil, nil
	}

	return result, nil
}

func (dao *DeliveryDaoImplGorm) SetState(delivery *types.Delivery, state types.DeliveryState) error {
	if state != types.DeliveryStateProgress &&
		state != types.DeliveryStateSuccess &&
		state != types.DeliveryStateFail {
		return errors.New("Wrong delivery state")
	}

	return dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&database.Delivery{}).
			Where("bot_id = ? AND campaign_id = ? AND telegram_id = ?",
				delivery.BotID,
				delivery.CampaignID,
				delivery.TelegramID).
			Update("state", state).Error; err != nil {
			return err
		}
		if state == types.DeliveryStateProgress {
			return nil
		}
		if err := dao.updateBotPossiblyEmptyStatus(tx, delivery.BotID, state == types.DeliveryStateFail); err != nil {
			return err
		}
		return nil
	})
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

func (dao *DeliveryDaoImplGorm) updateBotPossiblyEmptyStatus(tx *gorm.DB, botID int64, isPossiblyEmpty bool) error {
	if err := tx.
		Table("bots").
		Where("id = ?", botID).
		Updates(
			map[string]interface{}{
				"rr_possibly_empty": isPossiblyEmpty,
				"rr_access_time":    time.Now(),
			},
		).Error; err != nil {
		return err
	}
	return nil
}
