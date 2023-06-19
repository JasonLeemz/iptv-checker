package controllers

import (
	"iptv-checker/core/context"
	"iptv-checker/core/errors"
	"iptv-checker/internal/app/services"
)

type GetUserSourceReq struct {
	UserID int64 `json:"user_id"`
}

func GetUserSource(ctx ctx.Context) {
	var req GetUserSourceReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.Reply(nil, errors.GenErr(err))
		return
	}

	// Service 实例
	usService := services.NewUserSourceService()
	usDTO, err := usService.GetUserSource(req.UserID)

	ctx.Reply(usDTO, errors.GenErr(err), errors.ErrNoSuccess)
	//ctx.ViewData("message", usDTO)
	//ctx.View("test.html")

	//ctx.Writef("%v,%v", usDTO, err)
}
