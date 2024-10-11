package repositories

import (
	"context"
	"po/database"
	"po/grpc/user-grpc/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositories interface {
	CreateUser(ctx context.Context, model *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, model *models.User) (*models.User, error)
	PspHistory(ctx context.Context, id uuid.UUID) (*models.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (UserRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) CreateUser(ctx context.Context, model *models.User) (*models.User, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	model := &models.User{}

	if err := m.First(model, id).Error; err != nil {
		return nil, err
	}

	return model, nil

}

func (m *dbmanager) UpdateUser(ctx context.Context, model *models.User) (*models.User, error) {
	if err := m.Updates(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) PspHistory(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return nil, nil
}
