package types

type DeliveryState int

const (
	DeliveryStateProgress DeliveryState = 1
	DeliveryStateSuccess                = 2
	DeliveryStateFail                   = 3
)

type Delivery struct {
	CampaignID int64
	BotID      int64
	TelegramID int64
	State      DeliveryState
}
