package controllers

import (
	errors2 "errors"
	ctx "iptv-checker/core/context"
	"iptv-checker/core/errors"
	"iptv-checker/internal/app/logic"
	"iptv-checker/internal/app/services"
	"strings"
)

type AddSourceReq struct {
	UserID     int64  `json:"user_id"`
	SourceName string `json:"source_name"`
	SourceUrl  string `json:"source_url"`
}

type DelSourceReq struct {
	UserID       int64  `json:"user_id"`
	UserSourceID string `json:"user_source_id"`
	LiveSourceID string `json:"live_source_id"`
}

func AddSource(ctx ctx.Context) {
	// 创建 Service 实例
	usService := services.NewUserSourceService()

	// 使用 Service 获取用户信息
	var req AddSourceReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.Reply(nil, errors.GenErr(err))
		return
	}

	_, uAffected, err := usService.AddSource(req.UserID, req.SourceName, req.SourceUrl)
	ctx.Reply(uAffected, errors.GenErr(err))
}

func DelSource(ctx ctx.Context) {
	// 使用 Service 获取用户信息
	var req DelSourceReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.Reply(nil, errors.GenErr(err))
		return
	}
	userSourceID := strings.Split(req.UserSourceID, ",")
	liveSourceID := strings.Split(req.LiveSourceID, ",")

	// 校验是否能够删除
	if can := logic.CheckCanDelSource(req.UserID, userSourceID); can != true {
		ctx.Reply(nil, errors.GenErr(errors2.New("please check input")))
		return
	}

	// 创建 Service 实例
	usService := services.NewUserSourceService()

	err := usService.DelSource(userSourceID, liveSourceID)
	ctx.Reply(nil, errors.GenErr(err))
}
