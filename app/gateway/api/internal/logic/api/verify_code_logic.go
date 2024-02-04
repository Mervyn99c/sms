package api

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeLogic {
	return &VerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeLogic) VerifyCode(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
