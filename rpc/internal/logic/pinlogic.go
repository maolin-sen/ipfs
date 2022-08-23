package logic

import (
	"context"
	"fmt"
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

type PinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ipfs   iface.CoreAPI
	logx.Logger
}

func NewPinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PinLogic {
	addr, err := ma.NewMultiaddr(strings.TrimSpace(svcCtx.Config.IPFSAddr))
	if err != nil {
		return nil
	}
	ipfs, err := httpapi.NewApi(addr)
	if err != nil {
		return nil
	}
	return &PinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		ipfs:   ipfs,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PinLogic) Pin(in *ipfs.PinRequest) (*ipfs.PinResponse, error) {

	cid, err := cid.Parse(in.Cid.Str)
	if err != nil {
		return nil, err
	}
	path := path.IpfsPath(cid)

	s, err := l.ipfs.Object().Stat(l.ctx, path)
	if err != nil {
		return nil, fmt.Errorf("getting stats of cid %s:%s", cid, err)
	}

	return &ipfs.PinResponse{CumulativeSize: int32(s.CumulativeSize)}, nil
}
