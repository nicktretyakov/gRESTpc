package handlers

import (
	"context"
	"po/pb"
    "po/grpc/user-grpc/models"
	"po/grpc/user-grpc/repositories"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userRepositories repositories.UserRepositories
}

func NewuserHandler(userRepositories repositories.UserRepositories) (*UserHandler, error) {
	return &UserHandler{
		userRepositories: userRepositories,
	}, nil
}

func (h *UserHandler) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := &models.User{}
	err := copier.Copy(&user, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.userRepositories.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	pRes := &pb.User{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *UserHandler) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := &models.User{}
	err := copier.Copy(&user, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.userRepositories.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	pRes := &pb.User{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *UserHandler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Empty, error) {
	user := &models.User{}
	err := copier.Copy(&user, &req)
	if err != nil {
		return nil, err
	}

	id, _ := uuid.FromBytes([]byte(req.Id))
	res, err := h.userRepositories.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pRes := &pb.User{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (h *UserHandler) PspHistory(context.Context, *pb.PspHistoryRequest) (*pb.PspHistoryResponse, error) {

	return nil, nil
}
