package handlers

import (
	"context"
	"po/grpc/psp-grpc/models"
	"po/grpc/psp-grpc/repositories"
	"po/pb"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type PspHandler struct {
	pb.UnimplementedPspServiceServer
	pspRepositories repositories.PspRepositories
}

func NewPspHandler(jobRepository repositories.PspRepositories) (*PspHandler, error) {
	return &PspHandler{
		pspRepositories: jobRepository,
	}, nil
}

func (h *PspHandler) CreatePsp(ctx context.Context, req *pb.PspRequest) (*pb.Psp, error) {
	psp := &models.Psp{}
	err := copier.Copy(&psp, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.pspRepositories.CreatePsp(ctx, psp)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Psp{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *PspHandler) ViewPsp(ctx context.Context, req *pb.ViewPspRequest) (*pb.Psp, error) {
	psp := &models.Psp{}
	err := copier.Copy(&psp, &req)
	if err != nil {
		return nil, err
	}

	id, _ := uuid.FromBytes([]byte(req.PspId))
	res, err := h.pspRepositories.ViewPsp(ctx, id)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Psp{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *PspHandler) CancelPsp(ctx context.Context, req *pb.CancelPspRequest) (*pb.Empty, error) {
	psp := &models.Psp{}
	err := copier.Copy(&psp, &req)
	if err != nil {
		return nil, err
	}

	id, _ := uuid.FromBytes([]byte(req.PspId))
	_, err = h.pspRepositories.CancelPsp(ctx, id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
