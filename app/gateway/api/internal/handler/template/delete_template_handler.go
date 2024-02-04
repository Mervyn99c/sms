package template

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sms/app/gateway/api/internal/logic/template"
	"sms/app/gateway/api/internal/svc"
)

func DeleteTemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := template.NewDeleteTemplateLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTemplate()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
