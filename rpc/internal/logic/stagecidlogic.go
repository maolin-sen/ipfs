package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/internal/svc"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/ipfs"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/zeromicro/go-zero/core/logx"
)

type StageCidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ipfs   iface.CoreAPI
	logx.Logger
}

func NewStageCidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StageCidLogic {
	addr, err := ma.NewMultiaddr(strings.TrimSpace(svcCtx.Config.IPFSAddr))
	if err != nil {
		return nil
	}
	ipfs, err := httpapi.NewApi(addr)
	if err != nil {
		return nil
	}
	return &StageCidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		ipfs:   ipfs,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StageCidLogic) StageCid(in *ipfs.StageCidRequest) (*ipfs.Empty, error) {

	cid, err := cid.Parse(in.Cid.Str)
	if err != nil {
		return nil, err
	}
	if err := l.ipfs.Pin().Add(l.ctx, path.IpfsPath(cid), options.Pin.Recursive(true)); err != nil {
		return nil, fmt.Errorf("adding data to ipfs: %s", err)
	}

	return &ipfs.Empty{}, nil
}
