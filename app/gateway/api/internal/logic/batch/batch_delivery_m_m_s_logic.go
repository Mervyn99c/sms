package batch

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeliveryMMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchDeliveryMMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeliveryMMSLogic {
	return &BatchDeliveryMMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeliveryMMSLogic) BatchDeliveryMMS(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
