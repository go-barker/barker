package dao

import "github.com/my1562/userprofile/pkg/types"

type CampaignDao interface {
	Create(campaign *types.Campaign) (*types.Campaign, error)
	Update(campaign *types.Campaign) (*types.Campaign, error)
	Get(ID int64) (*types.Campaign, error)
	List() ([]types.Campaign, error)
}
