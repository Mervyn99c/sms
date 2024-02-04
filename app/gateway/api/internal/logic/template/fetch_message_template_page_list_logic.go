package template

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchMessageTemplatePageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchMessageTemplatePageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchMessageTemplatePageListLogic {
	return &FetchMessageTemplatePageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchMessageTemplatePageListLogic) FetchMessageTemplatePageList() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
