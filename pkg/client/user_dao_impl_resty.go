package client

import (
	"strconv"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/go-resty/resty/v2"
)

type UserDaoImplResty struct {
	resty *resty.Client
}

func NewUserDaoImplResty(resty *resty.Client) dao.UserDao {
	return &UserDaoImplResty{
		resty: resty,
	}
}

func (dao *UserDaoImplResty) Put(user *types.User) (*types.User, error) {
	resultWrapper := &struct{ Data *types.User }{Data: &types.User{}}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetBody(user).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"BotID": strconv.FormatInt(user.BotID, 10),
		}).
		Put("/bot/{BotID}/user")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}

func (dao *UserDaoImplResty) Get(botID int64, telegramID int64) (*types.User, error) {
	resultWrapper := &struct{ Data *types.User }{Data: &types.User{}}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		SetPathParams(map[string]string{
			"BotID":      strconv.FormatInt(botID, 10),
			"TelegramID": strconv.FormatInt(telegramID, 10),
		}).
		Get("/bot/{BotID}/user/{TelegramID}")
	if err != nil {
		return nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, nil
}

func (dao *UserDaoImplResty) List(botID int64, pageRequest *types.PaginatorRequest) ([]types.User, *types.PaginatorResponse, error) {
	resultWrapper := &struct {
		Data   []types.User
		Paging *types.PaginatorResponse
	}{}
	res, err := dao.resty.R().
		SetError(&ErrorResponse{}).
		SetResult(resultWrapper).
		SetQueryParams(pageRequest.ToMap()).
		SetPathParams(map[string]string{
			"BotID": strconv.FormatInt(botID, 10),
		}).
		Get("/bot/{BotID}/user")
	if err != nil {
		return nil, nil, err
	}
	if httpErr := res.Error(); httpErr != nil {
		return nil, nil, httpErr.(*ErrorResponse)
	}
	return resultWrapper.Data, resultWrapper.Paging, nil
}
