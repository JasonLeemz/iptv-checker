package ctx

import (
	"github.com/kataras/iris/v12"
	"iptv-checker/core/errors"
	"iptv-checker/core/log"
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
func (c *Context) Reply(obj interface{}, err *errors.Error, codes ...int32) {
	r := &reply{
		Data: obj,
	}

	code := int32(0)
	if len(codes) != 0 {
		code = codes[0]
		r.ErrNo = code
	}

	if err != nil {
		r.Msg = err.Error()
	} else {
		r.Msg = errors.GetErrMsg(code)
	}

	err2 := c.JSON(r)
	if err2 != nil {
		log.Logger.Error(err2)
	}
}
