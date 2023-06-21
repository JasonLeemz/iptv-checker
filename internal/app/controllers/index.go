package controllers

import (
	ctx "iptv-checker/core/context"
)

func Index(ctx ctx.Context) {

	ctx.View("index.html")

	//ctx.Writef("%v,%v", usDTO, err)
}

func Menu(ctx ctx.Context) {

	ctx.Reply(nil, nil)
}
func Demo(ctx ctx.Context) {

	ctx.View("demo.html")

	//ctx.Writef("%v,%v", usDTO, err)
}
