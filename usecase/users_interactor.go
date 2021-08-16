package usecase

import "github.com/kotaroooo0/lambda-clean-sample/entity"

type UsersInteractor interface {
	GetRecommendedUsers() ([]entity.User, error)
}

type UsersInteractorImpl struct {
	usersRepository UsersRepository
}

func NewUsersInteractorImpl(ur UsersRepository) *UsersInteractorImpl {
	return &UsersInteractorImpl{
		usersRepository: ur,
	}
}

func (ui *UsersInteractorImpl) GetRecommendedUsers() ([]entity.User, error) {
	return ui.usersRepository.GetRecommendedUsers()
}
