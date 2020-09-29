package dao

import (
	"github.com/corporateanon/barker/pkg/types"
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

func (dao *DeliveryDaoImplGorm) Take(campaignID int64) (*types.Delivery, *types.Campaign, *types.User, error) {
	panic("not implemented")
	// errNoRecepients := errors.New("no recepients")

	// resultingDelivery := &types.Delivery{}
	// var resultingUser *types.User

	// resultingCampaign, err := dao.campaignDao.Get(campaignID)
	// if err != nil {
	// 	return nil, nil, nil, err
	// }
	// if !resultingCampaign.Active {
	// 	return nil, nil, nil, errors.New("campaign not found or not active")
	// }

	// err = dao.db.Transaction(func(tx *gorm.DB) error {
	// 	userIDWrapper := &struct{ ID int64 }{}
	// 	query := tx.
	// 		Table("users").
	// 		Select("users.id").
	// 		Joins(
	// 			"left outer join deliveries on deliveries.user_id = users.id and deliveries.campaign_id = ?",
	// 			campaignID,
	// 		).
	// 		Where("deliveries.user_id IS NULL").
	// 		Where("users.deleted_at IS NULL").
	// 		Limit(1).
	// 		Scan(userIDWrapper)

	// 	if err := query.Error; err != nil {
	// 		return err
	// 	}
	// 	if userIDWrapper.ID == 0 {
	// 		return errNoRecepients
	// 	}

	// 	deliveryModel := &database.Delivery{
	// 		CampaignID: campaignID,
	// 		State:      types.DeliveryStateProgress,
	// 		UserID:     userIDWrapper.ID,
	// 	}
	// 	if err := tx.Create(deliveryModel).Error; err != nil {
	// 		return err
	// 	}

	// 	copier.Copy(resultingDelivery, deliveryModel)

	// 	resultingUser, err = dao.userDao.Get(resultingDelivery.UserID)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}

	// 	return nil
	// })

	// if err != nil {
	// 	if errors.Is(err, errNoRecepients) {
	// 		return nil, nil, nil, nil
	// 	}
	// 	return nil, nil, nil, err
	// }

	// return resultingDelivery, resultingCampaign, resultingUser, nil
}

func (dao *DeliveryDaoImplGorm) Success(campaignID int64, userID int64) error {
	panic("not implemented")
}

func (dao *DeliveryDaoImplGorm) Fail(campaignID int64, userID int64) error {
	panic("not implemented")
}
