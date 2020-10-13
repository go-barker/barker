package types

import (
	"errors"
	"strings"
)

type DeliveryState int

const (
	DeliveryStateProgress DeliveryState = 1
	DeliveryStateSuccess                = 2
	DeliveryStateFail                   = 3
)

func (state DeliveryState) ToString() (string, error) {
	switch state {
	case DeliveryStateProgress:
		return "Progress", nil
	case DeliveryStateSuccess:
		return "Success", nil
	case DeliveryStateFail:
		return "Fail", nil
	default:
		return "", errors.New("Wrong state")
	}
}

func DeliveryStateFromString(in string) (DeliveryState, error) {
	switch strings.ToLower(in) {
	case "progress":
		return DeliveryStateProgress, nil
	case "success":
		return DeliveryStateSuccess, nil
	case "fail":
		return DeliveryStateFail, nil
	default:
		return 0, errors.New("Wrong state")
	}
}

var AllDeliveryStates = []struct {
	Value  DeliveryState
	TSName string
}{
	{DeliveryStateProgress, "progress"},
	{DeliveryStateFail, "fail"},
	{DeliveryStateSuccess, "success"},
}

type Delivery struct {
	CampaignID int64         `json:"CampaignID,omitempty"`
	BotID      int64         `json:"BotID,omitempty"`
	TelegramID int64         `json:"TelegramID,omitempty"`
	State      DeliveryState `json:"State,omitempty"`
}
