package grpc

import (
	"context"

	"github.com/labstack/gommon/email"
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

func (h Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponce, error) {
	user, err := h.svc.CreateUser(req.Email)
	if err != nil {
		return &userpb.CreateUserResponce, err
	}
	return &userpb.CreateUserResponce{User: userpb.User{Id: user.Id.String(), Email: user.Email}}, nil
}
