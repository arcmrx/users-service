package grpc

import (
	"context"
	"fmt"

	userpb "github.com/arcmrx/project-protos/proto/user"
	"github.com/arcmrx/users-service/internal/user"
	"github.com/google/uuid"
)

type Handler struct {
	svc user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	email := req.Email
	user, err := h.svc.CreateUser(email)
	if err != nil {
		return &userpb.CreateUserResponse{}, err
	}

	pbUser := userpb.User{
		Id:    user.Id.String(),
		Email: user.Email,
	}
	return &userpb.CreateUserResponse{User: &pbUser}, nil
}
func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId, err := uuid.Parse(req.Id)
	if err != nil {
		return &userpb.GetUserResponse{}, err
	}

	user, err := h.svc.GetUser(userId)
	if err != nil {
		return &userpb.GetUserResponse{}, err
	}

	pbUser := &userpb.User{
		Id:    user.Id.String(),
		Email: user.Email,
	}

	return &userpb.GetUserResponse{
		User: pbUser,
	}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.ListUsers()
	if err != nil {
		return &userpb.ListUsersResponse{}, err
	}

	pbUsers := make([]*userpb.User, 0, len(users))
	for _, user := range users {
		usr := userpb.User{
			Id:    user.Id.String(),
			Email: user.Email,
		}
		pbUsers = append(pbUsers, &usr)
	}

	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	userId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID: %w, ID: %q", err, req.Id)
	}

	existingUser, err := h.svc.GetUser(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if req.Email != "" {
		existingUser.Email = req.Email
	}

	updatedUser, err := h.svc.UpdateUser(userId, existingUser.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    updatedUser.Id.String(),
			Email: updatedUser.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	idStr := req.Id
	userId, err := uuid.Parse(idStr)
	if err != nil {
		return &userpb.DeleteUserResponse{}, err
	}

	err = h.svc.DeleteUser(userId)
	if err != nil {
		return &userpb.DeleteUserResponse{}, err
	}
	return &userpb.DeleteUserResponse{}, nil
}
