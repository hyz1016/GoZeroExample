// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"GoZeroExample/application/article/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: ArticleDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: ArticleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/publish",
				Handler: PublishHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload/cover",
				Handler: UploadCoverHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1/article"),
	)
}
