package report

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WeMessageCallbackReceiverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWeMessageCallbackReceiverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WeMessageCallbackReceiverLogic {
	return &WeMessageCallbackReceiverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WeMessageCallbackReceiverLogic) WeMessageCallbackReceiver(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
