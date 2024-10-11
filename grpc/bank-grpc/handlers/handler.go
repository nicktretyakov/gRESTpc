package handlers

import (
	"context"
	"database/sql"
	"po/grpc/bank-grpc/requests"
	"po/grpc/bank-grpc/repositories"
	"po/grpc/bank-grpc/models"
	"po/pb"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BankHandler struct {
	pb.UnimplementedBankServiceServer
	bankRepository repositories.BankRepositories
}

func NewBankHandler(bankRepositories repositories.BankRepositories) (*BankHandler, error) {
	return &BankHandler{
		bankRepository: bankRepositories,
	}, nil
}

func (h *BankHandler) Create(ctx context.Context, req *pb.Bank) (*pb.Bank, error) {
	bank := &models.Bank{}
	err := copier.Copy(&bank, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.bankRepository.CreateBank(ctx, bank)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Bank{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *BankHandler) Update(ctx context.Context, req *pb.Bank) (*pb.Bank, error) {
	bank := &models.Bank{}
	err := copier.Copy(&bank, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.bankRepository.UpdateBank(ctx, bank)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Bank{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *BankHandler) List(ctx context.Context, req *pb.ListBankRequest) (*pb.ListBankResponse, error) {
	listBank := &requests.SearchBankRequest{}

	err := copier.Copy(&listBank, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.bankRepository.SearchBank(ctx, listBank)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "bank not found")
		}
		return nil, err
	}

	pRes := &pb.ListBankResponse{}
	err = copier.Copy(&pRes.Banks, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
