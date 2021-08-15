package usecase

import domain "github.com/kotaroooo0/lambda-clean-sample/domain/entity"

type UsersRepository interface {
	Get(userID int) (domain.User, error)
}
