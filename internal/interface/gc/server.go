package gc

import (
	"context"
	"github.com/DimKa163/graph/internal/core"
	"github.com/DimKa163/graph/internal/interface/gc/proto"
	"github.com/DimKa163/graph/internal/usecase"
	"github.com/beevik/guid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PathServer struct {
	appService *usecase.PathService
	proto.UnimplementedPathServiceServer
}

func NewPathServer(appService *usecase.PathService) *PathServer {
	return &PathServer{
		appService: appService,
	}
}

func (ps *PathServer) Register(server *grpc.Server) {
	proto.RegisterPathServiceServer(server, ps)
}

func (ps *PathServer) Get(ctx context.Context, in *proto.GetPath) (*proto.Path, error) {
	var path proto.Path
	if !in.HasId() {
		return nil, status.Error(codes.InvalidArgument, "missing path id")
	}
	id, err := guid.ParseString(in.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if !in.HasDefaultWarehouseId() {
		return nil, status.Error(codes.InvalidArgument, "missing default warehouse id")
	}
	defaultWhId, err := guid.ParseString(in.GetDefaultWarehouseId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	p, err := ps.appService.GetPath(ctx, id, defaultWhId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	list := p.GetList()
	nodes := make([]*proto.Node, len(list))
	for i, it := range list {
		node := &proto.Node{}
		node.SetId(it.ID.String())
		node.SetName(it.Name)
		node.SetType(mapType(it))
		node.SetTimeZone(it.Code)
		node.SetLevel(int32(it.Level))
		node.SetAvailableRest(it.AvailableRest)
		node.SetAddress(it.Address)
		node.SetOnlyStockPickupAllowed(it.OnlyStockPickupAllowed)
		nodes[i] = node
	}
	path.SetNode(nodes)
	return &path, nil
}

func mapType(n *core.PathNode) proto.NodeType {
	switch n.Type {
	case core.WarehouseMAIN:
		return proto.NodeType_MAIN
	case core.WarehouseCENTER:
		return proto.NodeType_CENTRAL
	case core.WarehouseFREE:
		return proto.NodeType_FREE
	case core.WarehouseMALL:
		return proto.NodeType_MALL
	default:
		return proto.NodeType_UNRECOGNIZED
	}
}
