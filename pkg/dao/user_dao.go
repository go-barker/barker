package dao

import "github.com/corporateanon/barker/pkg/types"

type UserDao interface {
	Put(user *types.User) (*types.User, error)
	Get(ID int64) (*types.User, error)
}
