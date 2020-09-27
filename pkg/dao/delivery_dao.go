package dao

import "github.com/my1562/userprofile/pkg/types"

type DeliveryDao interface {
	Take(campaignID int64) (*types.Delivery, *types.Campaign, *types.User, error)
	Success(campaignID int64, userID int64) error
	Fail(campaignID int64, userID int64) error
}
