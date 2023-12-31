// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cloud-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: UserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mail/code/send",
				Handler: MailCodeSendRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/share/basic/detail",
				Handler: ShareBasicDetailHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: FileUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/repository/save",
					Handler: UserRepositorySaveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/list",
					Handler: UserFilelistHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/name/update",
					Handler: UserFileNameUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/folder/create",
					Handler: UserFolderCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/file/delete",
					Handler: UserFileDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/file/move",
					Handler: UserFileMoveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/basic/create",
					Handler: ShareBasicCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/basic/save",
					Handler: ShareBasicSaveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload/prepare",
					Handler: FileUploadPrepareHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload/chunk",
					Handler: FileUploadChunkHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload/complete",
					Handler: FileUploadCompleteHandler(serverCtx),
				},
			}...,
		),
	)
}
