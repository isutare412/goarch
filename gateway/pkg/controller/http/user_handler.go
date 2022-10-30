package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/goarch/gateway/pkg/core/dto"
	"github.com/isutare412/goarch/gateway/pkg/core/port"
	"github.com/isutare412/goarch/gateway/pkg/log"
)

type userHandler struct {
	accSvc port.AccountService
}

func (h *userHandler) registerRoutes(g *gin.RouterGroup) {
	g.GET("/users/:nickname", h.getUser)
	g.POST("/users", h.createUser)
}

func (h *userHandler) getUser(c *gin.Context) {
	ctx := c.Request.Context()

	var pathParams pathParameters
	if err := c.ShouldBindUri(&pathParams); err != nil {
		log.WithOperation("shouldBindUri").Error(err)
		responseError(c, http.StatusBadRequest, err)
		return
	} else if err := pathParams.checkNickname(); err != nil {
		log.WithOperation("checkPathParams").Error(err)
		responseError(c, http.StatusBadRequest, err)
		return
	}

	var req = dto.GetUserByNicknameRequest{Nickname: pathParams.Nickname}
	var resp dto.GetUserByNicknameResponse
	resp, err := h.accSvc.GetUserByNickname(ctx, req)
	if err != nil {
		log.WithOperation("getUserByNickname").Error(err)
		responseError(c, http.StatusInternalServerError, err)
		return
	}
	responseJSON(c, http.StatusOK, resp)
}

func (h *userHandler) createUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithOperation("shouldBindJSON").Error(err)
		responseError(c, http.StatusBadRequest, err)
		return
	}

	if err := h.accSvc.CreateUser(ctx, req); err != nil {
		log.WithOperation("createUser").Error(err)
		responseError(c, http.StatusInternalServerError, err)
		return
	}
	responseStatus(c, http.StatusCreated)
}
