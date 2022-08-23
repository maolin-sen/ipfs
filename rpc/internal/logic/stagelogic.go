package logic

import (
	"context"
	"fmt"
	"io"
	"strings"

	ipfsfiles "github.com/ipfs/go-ipfs-files"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/internal/svc"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/ipfs"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/zeromicro/go-zero/core/logx"
)

type StageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ipfs   iface.CoreAPI
	logx.Logger
}

func NewStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StageLogic {
	addr, err := ma.NewMultiaddr(strings.TrimSpace(svcCtx.Config.IPFSAddr))
	if err != nil {
		return nil
	}
	ipfs, err := httpapi.NewApi(addr)
	if err != nil {
		return nil
	}
	return &StageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		ipfs:   ipfs,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StageLogic) Stage(in *ipfs.APIID) (*ipfs.Cid, error) {
	var r io.Reader
	path, err := l.ipfs.Unixfs().Add(l.ctx, ipfsfiles.NewReaderFile(r), options.Unixfs.Pin(true))
	if err != nil {
		return nil, fmt.Errorf("loading pinstore: %s", err)
	}

	return &ipfs.Cid{Str: path.Cid().String()}, nil
}
