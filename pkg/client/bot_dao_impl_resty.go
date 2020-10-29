package client

import (
	"strconv"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/go-resty/resty/v2"
)

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
		SetError(&ErrorResponse{}).
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
		SetError(&ErrorResponse{}).
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
		SetError(&ErrorResponse{}).
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

func (dao *BotDaoImplResty) List(pageRequest *types.PaginatorRequest) ([]types.Bot, *types.PaginatorResponse, error) {
	resultWrapper := &struct {
		Data   []types.Bot
		Paging *types.PaginatorResponse
	}{}
	pageRequestMap := pageRequest.ToMap()
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		SetQueryParams(pageRequestMap).
		Get("/bot")
	if err != nil {
		return nil, nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, resultWrapper.Paging, nil
}

func (dao *BotDaoImplResty) RRTake() (*types.Bot, error) {
	resultWrapper := &struct{ Data *types.Bot }{Data: &types.Bot{}}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		Post("/rr/bot")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}
