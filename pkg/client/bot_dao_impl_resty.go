package client

import (
	"strconv"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/go-resty/resty/v2"
)

type ErrorResponse struct {
	Message string `json:"error,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

type BotDaoImplResty struct {
	resty *resty.Client
}

func NewBotDaoImplResty(resty *resty.Client) dao.BotDao {
	return &BotDaoImplResty{
		resty: resty,
	}
}

func (dao *BotDaoImplResty) Create(bot *types.Bot) (*types.Bot, error) {
	resultWrapper := &struct{ Data *types.Bot }{Data: &types.Bot{}}
	res, err := dao.resty.R().
		SetBody(bot).
		SetResult(resultWrapper).
		Post("/bot")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}

func (dao *BotDaoImplResty) Update(bot *types.Bot) (*types.Bot, error) {
	resultWrapper := &struct{ Data *types.Bot }{Data: &types.Bot{}}
	res, err := dao.resty.R().
		SetBody(bot).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"id": strconv.FormatInt(bot.ID, 10),
		}).
		Put("/bot/{id}")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}

func (dao *BotDaoImplResty) Get(ID int64) (*types.Bot, error) {
	resultWrapper := &struct{ Data *types.Bot }{Data: &types.Bot{}}
	res, err := dao.resty.R().
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"id": strconv.FormatInt(ID, 10),
		}).
		Get("/bot/{id}")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}

func (dao *BotDaoImplResty) List() ([]types.Bot, error) {
	panic("not implemented")
}