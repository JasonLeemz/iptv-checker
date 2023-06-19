package controllers

import (
	"iptv-checker/core/context"
	"iptv-checker/core/errors"
	"iptv-checker/internal/app/services"
)

// ctx iris.Context

func GetUserSource(ctx ctx.Context) {
	// Service 实例
	usService := services.NewUserSourceService()

	// 使用 Service 获取用户信息
	userID := int64(5639017)
	usDTO, err := usService.GetUserSource(userID)

	ctx.Reply(usDTO, errors.GenErr(err), errors.ErrNoSuccess)
	//ctx.ViewData("message", usDTO)
	//ctx.View("test.html")

	//ctx.Writef("%v,%v", usDTO, err)
}
