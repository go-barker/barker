package types

import "time"

type Bot struct {
	ID              int64     `json:"ID,omitempty"`
	Title           string    `binding:"required" json:"Title,omitempty"`
	Token           string    `binding:"required" json:"Token,omitempty"`
	RRAccessTime    time.Time `json:"RRAccessTime,omitempty" ts_type:"string"`
	RRPossiblyEmpty bool      `json:"RRPossiblyEmpty,omitempty"`
}
