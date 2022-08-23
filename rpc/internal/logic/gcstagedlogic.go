package logic

import (
	"context"
	"strings"

	httpapi "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/internal/svc"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/ipfs"
)

type GCStagedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ipfs   iface.CoreAPI
	logx.Logger
}

func NewGCStagedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GCStagedLogic {
	addr, err := ma.NewMultiaddr(strings.TrimSpace(svcCtx.Config.IPFSAddr))
	if err != nil {
		return nil
	}
	ipfs, err := httpapi.NewApi(addr)
	if err != nil {
		return nil
	}
	return &GCStagedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		ipfs:   ipfs,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GCStagedLogic) GCStaged(in *ipfs.GCStagedRequest) (*ipfs.GCStagedResponse, error) {
	// todo: add your logic here and delete this line

	return &ipfs.GCStagedResponse{}, nil
}
