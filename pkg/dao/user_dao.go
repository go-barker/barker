package dao

import "github.com/my1562/userprofile/pkg/types"

type UserDao interface {
	Put(user *types.User) (*types.User, error)
	Get(ID int64) (*types.User, error)
}
