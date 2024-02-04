package template

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTemplateLogic {
	return &DeleteTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTemplateLogic) DeleteTemplate() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
