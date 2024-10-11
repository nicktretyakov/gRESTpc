package repositories

import (
	"context"
	"po/database"
	"po/grpc/bank-grpc/models"
	"po/grpc/bank-grpc/requests"

	"gorm.io/gorm"
)

type BankRepositories interface {
	CreateBank(ctx context.Context, model *models.Bank) (*models.Bank, error)
	UpdateBank(ctx context.Context, model *models.Bank) (*models.Bank, error)
	SearchBank(ctx context.Context, req *requests.SearchBankRequest) ([]*models.Bank, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BankRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Bank{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) CreateBank(ctx context.Context, model *models.Bank) (*models.Bank, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) UpdateBank(ctx context.Context, model *models.Bank) (*models.Bank, error) {
	bank := models.Bank{}
	if err := m.Updates(model).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}

func (m *dbmanager) SearchBank(ctx context.Context, req *requests.SearchBankRequest) ([]*models.Bank, error) {
	banks := []*models.Bank{}
	if req.Level > 0 {
		if err := m.Where("level > ?", req.Level).Find(&banks).Error; err != nil {
			return nil, err
		}
	} else {
		if err := m.Find(&banks).Error; err != nil {
			return nil, err
		}
	}

	return banks, nil
}
