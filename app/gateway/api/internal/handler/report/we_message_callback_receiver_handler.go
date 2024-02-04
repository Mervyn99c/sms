package report

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sms/app/gateway/api/internal/logic/report"
	"sms/app/gateway/api/internal/svc"
	"sms/app/gateway/api/internal/types"
)

func WeMessageCallbackReceiverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := report.NewWeMessageCallbackReceiverLogic(r.Context(), svcCtx)
		resp, err := l.WeMessageCallbackReceiver(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
