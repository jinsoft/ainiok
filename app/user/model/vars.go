package model

import "github.com/tal-tech/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	AuthByPassword = iota
	AuthByCode
	AuthByToken
)
