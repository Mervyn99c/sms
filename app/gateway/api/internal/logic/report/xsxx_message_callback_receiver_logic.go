package report

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type XsxxMessageCallbackReceiverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewXsxxMessageCallbackReceiverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *XsxxMessageCallbackReceiverLogic {
	return &XsxxMessageCallbackReceiverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *XsxxMessageCallbackReceiverLogic) XsxxMessageCallbackReceiver(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
