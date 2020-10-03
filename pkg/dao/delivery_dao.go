package dao

import "github.com/corporateanon/barker/pkg/types"

type DeliveryDao interface {
	Take(botID int64, campaignID int64) (*types.Delivery, *types.User, error)
	Success(campaignID int64, userID int64) error
	Fail(campaignID int64, userID int64) error
}
