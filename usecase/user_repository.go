package usecase

import (
	"github.com/yama-koo/go-api-example/domain"
)

// UserRepository interface
type UserRepository interface {
	Store(domain.User) (int, error)
	FindByID(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
