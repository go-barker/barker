package client

import (
	"strconv"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/go-resty/resty/v2"
)

type DeliveryDaoImplResty struct {
	resty *resty.Client
}

func NewDeliveryDaoImplResty(
	resty *resty.Client,
) dao.DeliveryDao {
	return &DeliveryDaoImplResty{
		resty: resty,
	}
}

func (dao *DeliveryDaoImplResty) Take(botID int64, campaignID int64) (*types.Delivery, *types.User, error) {
	resultWrapper := &struct {
		Data *types.Delivery
		User *types.User
	}{
		Data: &types.Delivery{},
		User: &types.User{},
	}

	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"BotID":      strconv.FormatInt(botID, 10),
			"CampaignID": strconv.FormatInt(campaignID, 10),
		}).
		Post("/bot/{BotID}/campaign/{CampaignID}/delivery")
	if err != nil {
		return nil, nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, resultWrapper.User, nil
}

func (dao *DeliveryDaoImplResty) SetState(delivery *types.Delivery, state types.DeliveryState) error {
	stateString, err := state.ToString()
	if err != nil {
		return err
	}

	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetPathParams(map[string]string{
			"BotID":      strconv.FormatInt(delivery.BotID, 10),
			"CampaignID": strconv.FormatInt(delivery.CampaignID, 10),
			"TelegramID": strconv.FormatInt(delivery.TelegramID, 10),
			"State":      stateString,
		}).
		Put("/bot/{BotID}/campaign/{CampaignID}/delivery/{TelegramID}/state/{State}")
	if err != nil {
		return err
	}
	if httpErr := res.Error(); httpErr != nil {
		return httpErr.(*ErrorResponse)
	}
	return nil
}

func (dao *DeliveryDaoImplResty) GetState(delivery *types.Delivery) (types.DeliveryState, error) {
	resultWrapper := &struct {
		Data string
	}{}

	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"BotID":      strconv.FormatInt(delivery.BotID, 10),
			"CampaignID": strconv.FormatInt(delivery.CampaignID, 10),
			"TelegramID": strconv.FormatInt(delivery.TelegramID, 10),
		}).
		Get("/bot/{BotID}/campaign/{CampaignID}/delivery/{TelegramID}/state")
	if err != nil {
		return 0, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return 0, httpErr.(*ErrorResponse)
	}
	state, err := types.DeliveryStateFromString(resultWrapper.Data)
	if err != nil {
		return 0, err
	}
	return state, nil
}
