package controllers

import (
	ctx "iptv-checker/core/context"
)

func Index(ctx ctx.Context) {

	ctx.View("index.html")

	//ctx.Writef("%v,%v", usDTO, err)
}
