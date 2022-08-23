// Code generated by goctl. DO NOT EDIT!
// Source: storage.proto

package server

import (
	"context"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/internal/svc"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/ipfs"
	"github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/internal/logic"
)

type IPFSServer struct {
	svcCtx *svc.ServiceContext
	ipfs.UnimplementedIPFSServer
}

func NewIPFSServer(svcCtx *svc.ServiceContext) *IPFSServer {
	return &IPFSServer{
		svcCtx: svcCtx,
	}
}

func (s *IPFSServer) Stage(ctx context.Context, in *ipfs.APIID) (*ipfs.Cid, error) {
	l := logic.NewStageLogic(ctx, s.svcCtx)
	return l.Stage(in)
}

func (s *IPFSServer) StageCid(ctx context.Context, in *ipfs.StageCidRequest) (*ipfs.Empty, error) {
	l := logic.NewStageCidLogic(ctx, s.svcCtx)
	return l.StageCid(in)
}

func (s *IPFSServer) Unpin(ctx context.Context, in *ipfs.UnpinRequest) (*ipfs.Empty, error) {
	l := logic.NewUnpinLogic(ctx, s.svcCtx)
	return l.Unpin(in)
}

func (s *IPFSServer) Pin(ctx context.Context, in *ipfs.PinRequest) (*ipfs.PinResponse, error) {
	l := logic.NewPinLogic(ctx, s.svcCtx)
	return l.Pin(in)
}

func (s *IPFSServer) Get(ctx context.Context, in *ipfs.Cid) (*ipfs.Empty, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}

func (s *IPFSServer) Replace(ctx context.Context, in *ipfs.ReplaceRequest) (*ipfs.ReplaceResponse, error) {
	l := logic.NewReplaceLogic(ctx, s.svcCtx)
	return l.Replace(in)
}

func (s *IPFSServer) IsPinned(ctx context.Context, in *ipfs.IsPinnedRequest) (*ipfs.IsPinnedResponse, error) {
	l := logic.NewIsPinnedLogic(ctx, s.svcCtx)
	return l.IsPinned(in)
}

func (s *IPFSServer) GCStaged(ctx context.Context, in *ipfs.GCStagedRequest) (*ipfs.GCStagedResponse, error) {
	l := logic.NewGCStagedLogic(ctx, s.svcCtx)
	return l.GCStaged(in)
}

func (s *IPFSServer) PinnedCids(ctx context.Context, in *ipfs.Empty) (*ipfs.Empty, error) {
	l := logic.NewPinnedCidsLogic(ctx, s.svcCtx)
	return l.PinnedCids(in)
}