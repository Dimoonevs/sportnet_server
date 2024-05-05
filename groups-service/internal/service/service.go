package service

import (
	"context"

	"github.com/Dimoonevs/SportAspp/groups-service/internal/repository"
	proto "github.com/Dimoonevs/SportAspp/groups-service/proto/groups"
)

type GroupService struct {
	proto.UnimplementedGroupServiceServer
	Repo repository.GroupRepository
}

func (g *GroupService) CreateGroup(ctx context.Context, req *proto.GroupRequest) (*proto.GroupResponse, error) {

	id, err := g.Repo.CreateGroup(ctx, req)
	if err != nil {
		return nil, err
	}

	return &proto.GroupResponse{
		Message: "Group created successfully",
		Id:      id,
	}, nil
}
func (g *GroupService) GetGroups(ctx context.Context, req *proto.GetGroupRequest) (*proto.GetGroupResponse, error) {
	groups, err := g.Repo.GetGroups(ctx, req)
	if err != nil {
		return nil, err
	}
	return &proto.GetGroupResponse{
		Groups: groups,
	}, nil
}

func (g *GroupService) EditGroup(ctx context.Context, req *proto.GroupEditRequest) (*proto.GroupResponse, error) {
	err := g.Repo.EditGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	err = g.Repo.ChangeGroupsSubscriptions(ctx, req.Id, req.SubscriptionId)
	if err != nil {
		return nil, err
	}
	return &proto.GroupResponse{
		Message: "Group edited successfully",
	}, nil
}
