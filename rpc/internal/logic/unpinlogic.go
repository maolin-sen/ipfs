package logic

import (
	"context"
	"strings"

	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/maolin-sen/ipfs/rpc/internal/svc"
	"github.com/maolin-sen/ipfs/rpc/ipfs"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/zeromicro/go-zero/core/logx"
)

type UnpinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ipfs   iface.CoreAPI
	logx.Logger
}

func NewUnpinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnpinLogic {
	addr, err := ma.NewMultiaddr(strings.TrimSpace(svcCtx.Config.IPFSAddr))
	if err != nil {
		return nil
	}
	ipfs, err := httpapi.NewApi(addr)
	if err != nil {
		return nil
	}
	return &UnpinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		ipfs:   ipfs,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnpinLogic) Unpin(in *ipfs.UnpinRequest) (*ipfs.Empty, error) {

	cid, err := cid.Parse(in.Cid.Str)
	if err != nil {
		return nil, err
	}

	if err := l.ipfs.Pin().Rm(l.ctx, path.IpfsPath(cid), options.Pin.RmRecursive(true)); err != nil {
		return nil, err
	}

	return &ipfs.Empty{}, nil
}
