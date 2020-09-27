package types

type DeliveryState int

const (
	DeliveryStateProgress DeliveryState = 1
	DeliveryStateSuccess                = 2
	DeliveryStateFail                   = 3
)

type Delivery struct {
	CampaignID int64
	UserID     int64
	State      DeliveryState
}
