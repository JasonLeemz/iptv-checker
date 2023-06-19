package router

import (
	"github.com/kataras/iris/v12/context"
	ctx "iptv-checker/core/context"
)

func Get(c *context.Context, fun func(ctx ctx.Context)) func(c *context.Context) {
	return func(c *context.Context) {

	}
}