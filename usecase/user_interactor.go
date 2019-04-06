package usecase

import (
	"github.com/yama-koo/go-api-example/domain"
)

// UserInteractor struct
type UserInteractor struct {
	UserRepository UserRepository
}

// Add func
func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

// Users func
func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

// FindByID func
func (interactor *UserInteractor) FindByID(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(id)
	return
}
