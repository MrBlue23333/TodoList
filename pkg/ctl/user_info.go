package ctl

import (
	"context"
	"errors"
)

var userKey int

type UserInfo struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("failed to get userinfo")
	}
	return user, nil
}

// NewContext calls context.With
func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}
