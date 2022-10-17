package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	// 增加点击次数
	_, err = l.svcCtx.Engine.Exec("update share_basic set click_num = click_num + 1 where identity = ?", req.Identity)
	if err != nil {
		return nil, err
	}
	// 获取详情
	resp = new(types.ShareBasicDetailReply)
	_, err = l.svcCtx.Engine.Table("share_basic").Where("share_basic.identity = ?", req.Identity).
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("left", "repository_pool", "repository_pool.identity = share_basic.repository_identity").
		Join("left", "user_repository", "user_repository.identity = share_basic.user_repository_identity").
		Get(resp)
	if err != nil {
		return nil, err
	}

	return
}
