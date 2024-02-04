package template

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchMessageTemplateAndConfigByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchMessageTemplateAndConfigByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchMessageTemplateAndConfigByIdLogic {
	return &FetchMessageTemplateAndConfigByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchMessageTemplateAndConfigByIdLogic) FetchMessageTemplateAndConfigById() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
