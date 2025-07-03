package grpc

import (
	"context"

	userpb "github.com/arcmrx/project-protos/proto/user"
	"github.com/arcmrx/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := h.svc.CreateUser(req.Email)
	if err != nil {
		return &userpb.CreateUserResponse{}, err
	}
	return &userpb.CreateUserResponse{User: &userpb.User{Id: user.Id.String(), Email: user.Email}}, nil
}
