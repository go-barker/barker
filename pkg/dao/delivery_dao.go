package dao

import "github.com/corporateanon/barker/pkg/types"

type DeliveryTakeResult struct {
	Delivery *types.Delivery
	Campaign *types.Campaign
	User     *types.User
}

type DeliveryDao interface {
	Take(botID int64, campaignID int64) (*DeliveryTakeResult, error)
	SetState(*types.Delivery, types.DeliveryState) error
	GetState(*types.Delivery) (types.DeliveryState, error)
}
