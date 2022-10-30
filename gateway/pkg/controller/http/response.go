package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

func responseError(c *gin.Context, code int, err error) {
	var errMsg = err.Error()
	if kerr := pkgerr.AsKnown(err); kerr != nil {
		errMsg = kerr.SimpleError()
		// TODO: Overwrite error code based on errno.
	}
	c.JSON(code, errorResponse{Msg: errMsg})
}

func responseStatus(c *gin.Context, code int) {
	c.Status(code)
}

func responseJSON(c *gin.Context, code int, resp any) {
	c.JSON(code, resp)
}
