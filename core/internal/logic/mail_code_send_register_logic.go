package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRegisterRequest) (resp *types.MailCodeSendRegisterReply, err error) {
	cnt, err := models.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	fmt.Println(cnt)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("邮箱已注册")
		return
	}
	code := helper.RandCode()
	models.RDB.Set(l.ctx, req.Email, code, define.CodeExpired)
	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
