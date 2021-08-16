package usecase

import "github.com/kotaroooo0/lambda-clean-sample/entity"

type UsersRepository interface {
	Get(userID int) (*entity.User, error)
	GetRecommendedUsers() ([]entity.User, error)
}
