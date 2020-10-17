package dao

import "github.com/corporateanon/barker/pkg/types"

type CampaignDao interface {
	Create(campaign *types.Campaign) (*types.Campaign, error)
	Update(campaign *types.Campaign) (*types.Campaign, error)
	Get(botID int64, ID int64) (*types.Campaign, error)
	List(botID int64, pageRequest *types.PaginatorRequest) ([]types.Campaign, *types.PaginatorResponse, error)
}
