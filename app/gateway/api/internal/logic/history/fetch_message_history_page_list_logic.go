package history

import (
	"context"

	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchMessageHistoryPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchMessageHistoryPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchMessageHistoryPageListLogic {
	return &FetchMessageHistoryPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchMessageHistoryPageListLogic) FetchMessageHistoryPageList(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
