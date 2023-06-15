package controllers

import (
	"github.com/kataras/iris/v12"
	"iptv-checker/internal/app/dao"
	"iptv-checker/internal/app/services"
)

func GetUserSource(ctx iris.Context) {
	// 创建 DAO 和 Service 实例
	usDAO := dao.NewUserSourceDAO()
	usService := services.NewUserSourceService(usDAO)

	// 使用 Service 获取用户信息
	userID := int64(5639017)
	usDTO, _ := usService.GetUserSource(userID)

	ctx.JSON(usDTO)
	//ctx.Writef("%v,%v", usDTO, err)
}
