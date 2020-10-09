package dao

import (
	"errors"

	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type CampaignDaoImplGorm struct {
	db *gorm.DB
}

func NewCampaignDaoImplGorm(db *gorm.DB) CampaignDao {
	return &CampaignDaoImplGorm{
		db: db,
	}
}

func (dao *CampaignDaoImplGorm) Create(campaign *types.Campaign) (*types.Campaign, error) {
	campaignModel := &database.Campaign{}
	campaignModel.FromEntity(campaign)
	if err := dao.db.Create(campaignModel).Error; err != nil {
		return nil, err
	}
	resultingCampaign := &types.Campaign{}
	campaignModel.ToEntity(resultingCampaign)
	return resultingCampaign, nil
}

func (dao *CampaignDaoImplGorm) Update(campaign *types.Campaign) (*types.Campaign, error) {
	if campaign.ID == 0 {
		return nil, errors.New("ID missing")
	}
	campaignModel := &database.Campaign{}

	if err := dao.db.
		Where("id = ? AND bot_id = ?", campaign.ID, campaign.BotID).
		First(campaignModel).Error; err != nil {
		return nil, err
	}

	campaignModel.FromEntity(campaign)

	if err := dao.db.Save(campaignModel).Error; err != nil {
		return nil, err
	}
	resultingCampaign := &types.Campaign{}
	campaignModel.ToEntity(resultingCampaign)
	return resultingCampaign, nil
}

func (dao *CampaignDaoImplGorm) Get(botID int64, ID int64) (*types.Campaign, error) {
	campaignModel := &database.Campaign{}

	if err := dao.db.
		Where("id = ?", ID).
		Where("bot_id = ?", botID).
		First(campaignModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	resultingCampaign := &types.Campaign{ID: ID}
	campaignModel.ToEntity(resultingCampaign)
	return resultingCampaign, nil
}

func (dao *CampaignDaoImplGorm) List() ([]types.Campaign, error) {
	panic("not implemented")
}
