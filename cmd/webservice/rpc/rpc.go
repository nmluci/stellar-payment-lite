package rpc

import (
	insvc "github.com/nmluci/go-backend/internal/service"
	hentaiRPC "github.com/nmluci/gostellar/pkg/rpc/hentai"
)

type HentaiRPC struct {
	hentaiRPC.UnimplementedNakaZettaiDameServer
	service insvc.Service
}

func Init(svc insvc.Service) hentaiRPC.NakaZettaiDameServer {
	return &HentaiRPC{
		service: svc,
	}
}
