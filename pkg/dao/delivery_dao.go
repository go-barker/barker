package dao

import "github.com/corporateanon/barker/pkg/types"

type DeliveryDao interface {
	Take(botID int64, campaignID int64) (*types.Delivery, *types.User, error)
	SetState(*types.Delivery, types.DeliveryState) error
	GetState(*types.Delivery) (types.DeliveryState, error)
}
