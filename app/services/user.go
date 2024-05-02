package services

import (
	"context"
	"dgut-icourse/ent"
	"dgut-icourse/ent/user"
	"fmt"
	"github.com/google/uuid"
)

// GetOrCreateUserByOpenID
// @Description: 通过openid获取用户信息，若用户不存在则创建用户
// @param ctx: context.Context
// @param openid: string
// @return *ent.User: 用户信息
// @return error: err
func GetOrCreateUserByOpenID(ctx context.Context, openid string) (u *ent.User, err error) {
	if u, err = Client.User.Query().Where(user.OpenidEQ(openid)).Only(ctx); err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		} else {
			u, err = Client.User.Create().SetOpenid(openid).Save(ctx)
		}
	}
	return u, nil
}

// GetUserByUUID
// @Description: 通过UID获取用户信息
// @param ctx: context.Context
// @param uid: string
// @return *ent.User: 用户信息
// @return error: err
func GetUserByUUID(ctx context.Context, uuidString string) (u *ent.User, err error) {
	if uid, err := uuid.Parse(uuidString); err != nil {
		return nil, err
	} else {
		return Client.User.Get(ctx, uid)
	}
}

// UpdateMe
// @Description: 更新用户信息
// @param ctx: context.Context
// @param user: ent.User
// @return *ent.User: 用户信息
// @return error: err
func UpdateMe(ctx context.Context, user ent.User) (u *ent.User, err error) {
	if u = ctx.Value("user").(*ent.User); u == nil {
		return nil, fmt.Errorf("invalid user")
	} else {
		return Client.User.UpdateOne(u).SetNickName(*user.NickName).SetIcon(*user.Icon).Save(ctx)
	}
}
