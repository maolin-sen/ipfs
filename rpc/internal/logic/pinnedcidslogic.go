package logic

import (
	"context"
	"strings"

	httpapi "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/maolin-sen/ipfs/rpc/internal/svc"
	"github.com/maolin-sen/ipfs/rpc/ipfs"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/zeromicro/go-zero/core/logx"
)

type PinnedCidsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ipfs   iface.CoreAPI
	logx.Logger
}

func NewPinnedCidsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PinnedCidsLogic {
	addr, err := ma.NewMultiaddr(strings.TrimSpace(svcCtx.Config.IPFSAddr))
	if err != nil {
		return nil
	}
	ipfs, err := httpapi.NewApi(addr)
	if err != nil {
		return nil
	}
	return &PinnedCidsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		ipfs:   ipfs,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PinnedCidsLogic) PinnedCids(in *ipfs.Empty) (*ipfs.Empty, error) {
	// todo: add your logic here and delete this line

	return &ipfs.Empty{}, nil
}
