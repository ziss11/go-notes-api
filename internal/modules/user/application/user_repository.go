package application

import "notes-api/internal/modules/user/domain"

type UserRepository interface {
	FindAll() ([]*domain.User, error)
	FindByID(id string) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(user *domain.User) error
}
