package logic

import (
	"cloud-disk/core/models"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	resp = &types.UserDetailReply{}
	ud := new(models.UserBasic)
	has, err := models.Engine.Where("identity=?", req.Identity).Get(ud)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("User Not Found")
	}
	resp.Name = ud.Name
	resp.Email = ud.Email
	return
}
