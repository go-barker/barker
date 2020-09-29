package dao

import "github.com/corporateanon/barker/pkg/types"

type UserDao interface {
	Put(user *types.User) (*types.User, error)
	Get(botID int64, telegramID int64) (*types.User, error)
}
