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

type ReplaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ipfs   iface.CoreAPI
	logx.Logger
}

func NewReplaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplaceLogic {
	addr, err := ma.NewMultiaddr(strings.TrimSpace(svcCtx.Config.IPFSAddr))
	if err != nil {
		return nil
	}
	ipfs, err := httpapi.NewApi(addr)
	if err != nil {
		return nil
	}
	return &ReplaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		ipfs:   ipfs,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReplaceLogic) Replace(in *ipfs.ReplaceRequest) (*ipfs.ReplaceResponse, error) {
	// todo: add your logic here and delete this line

	return &ipfs.ReplaceResponse{}, nil
}
