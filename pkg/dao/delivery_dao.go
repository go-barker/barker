package dao

import "github.com/corporateanon/barker/pkg/types"

type DeliveryTakeResult struct {
	Delivery *types.Delivery `json:"Delivery,omitempty"`
	Campaign *types.Campaign `json:"Campaign,omitempty"`
	User     *types.User     `json:"User,omitempty"`
}

type DeliveryDao interface {
	Take(botID int64, campaignID int64, telegramID int64) (*DeliveryTakeResult, error)
	SetState(*types.Delivery, types.DeliveryState) error
	GetState(*types.Delivery) (types.DeliveryState, error)
}
