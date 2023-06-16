package ctx

import (
	"github.com/kataras/iris/v12"
	"iptv-checker/pkg/errors"
)

type Context struct {
	iris.Context
}

type reply struct {
	ErrNo int32       `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

// Reply ...
func (c *Context) Reply(code int32, obj interface{}, err *errors.Error) {
	r := &reply{
		Data: obj,
	}
	if err != nil {
		r.Msg = err.Error()
	} else {
		r.Msg = errors.GetErrMsg(code)
	}

	c.JSON(r)
}
