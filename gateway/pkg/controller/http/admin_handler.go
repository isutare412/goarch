package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isutare412/goarch/gateway/pkg/core/port"
)

type adminHandler struct {
	accSvc port.AccountService
}

func (h *adminHandler) registerRoutes(g *gin.RouterGroup) {
}
