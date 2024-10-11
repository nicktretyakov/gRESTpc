package repositories

import (
	"context"
	"po/database"
	"po/grpc/psp-grpc/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PspRepositories interface {
	CreatePsp(ctx context.Context, model *models.Psp) (*models.Psp, error)
	ViewPsp(ctx context.Context, id uuid.UUID) (*models.Psp, error)
	CancelPsp(ctx context.Context, id uuid.UUID) (*models.Psp, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (PspRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Psp{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) CreatePsp(ctx context.Context, model *models.Psp) (*models.Psp, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) ViewPsp(ctx context.Context, id uuid.UUID) (*models.Psp, error) {
	psp := models.Psp{}
	if err := m.First(&psp, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	return &psp, nil
}

func (m *dbmanager) CancelPsp(ctx context.Context, id uuid.UUID) (*models.Psp, error) {
	psp := models.Psp{}
	if err := m.First(&psp, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	psp.Status = "Cancelled"
	if err := m.Updates(psp).Error; err != nil {
		return nil, err
	}
	return &psp, nil
}
