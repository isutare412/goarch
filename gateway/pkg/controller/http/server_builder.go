package http

import (
	"github.com/isutare412/goarch/gateway/pkg/config"
	"github.com/isutare412/goarch/gateway/pkg/core/port"
)

type serverBuilder struct {
	cfg config.HTTPServerConfig

	accSvc port.AccountService
	mtgSvc port.MeetingService
}

func ServerBuilder() serverBuilder { return serverBuilder{} }

func (b serverBuilder) WithHTTPServerConfig(cfg config.HTTPServerConfig) serverBuilder {
	b.cfg = cfg
	return b
}

func (b serverBuilder) WithAccountService(accSvc port.AccountService) serverBuilder {
	b.accSvc = accSvc
	return b
}

func (b serverBuilder) WithMeetingService(mtgSvc port.MeetingService) serverBuilder {
	b.mtgSvc = mtgSvc
	return b
}

func (b serverBuilder) Build() *Server {
	return &Server{
		cfg:    b.cfg,
		accSvc: b.accSvc,
		mtgSvc: b.mtgSvc,
	}
}
