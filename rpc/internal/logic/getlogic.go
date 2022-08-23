package logic

import (
	"context"
	"strings"

	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/maolin-sen/ipfs/rpc/internal/svc"
	"github.com/maolin-sen/ipfs/rpc/ipfs"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ipfs   iface.CoreAPI
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	addr, err := ma.NewMultiaddr(strings.TrimSpace(svcCtx.Config.IPFSAddr))
	if err != nil {
		return nil
	}
	ipfs, err := httpapi.NewApi(addr)
	if err != nil {
		return nil
	}
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		ipfs:   ipfs,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *ipfs.Cid) (*ipfs.Empty, error) {
	cid, err := cid.Parse(in.Str)
	if err != nil {
		return nil, err
	}

	node, err := l.ipfs.Unixfs().Get(l.ctx, path.IpfsPath(cid))

	//todo:
	println(node)

	return &ipfs.Empty{}, nil
}
