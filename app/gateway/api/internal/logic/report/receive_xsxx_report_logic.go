package report

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveXsxxReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceiveXsxxReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveXsxxReportLogic {
	return &ReceiveXsxxReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveXsxxReportLogic) ReceiveXsxxReport(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
