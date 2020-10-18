package dbclient

import (
	"errors"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/pagination"
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type CampaignDaoImplGorm struct {
	db *gorm.DB
}

func NewCampaignDaoImplGorm(db *gorm.DB) dao.CampaignDao {
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

func (dao *CampaignDaoImplGorm) List(botID int64, pageRequest *types.PaginatorRequest) ([]types.Campaign, *types.PaginatorResponse, error) {
	campaignModelsList := []database.Campaign{}
	db := dao.db.Table("campaigns").
		Order("created_at DESC").
		Where("bot_id = ?", botID)
	resp := pagination.Paging(&pagination.Param{
		DB:    db,
		Page:  int(pageRequest.Page),
		Limit: int(pageRequest.Size),
	}, &campaignModelsList)

	if err := db.Error; err != nil {
		return nil, nil, err
	}

	campaignsList := make([]types.Campaign, len(campaignModelsList))
	for i, model := range campaignModelsList {
		model.ToEntity(&campaignsList[i])
	}
	return campaignsList,
		&types.PaginatorResponse{
			Page:       resp.Page,
			Size:       resp.Limit,
			Total:      resp.TotalPage,
			TotalItems: resp.TotalRecord,
		},
		nil
}
