package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kotaroooo0/lambda-clean-sample/entity"
	"github.com/kotaroooo0/lambda-clean-sample/usecase"
)

type UsersController struct {
	usersInteractor usecase.UsersInteractor
}

func NewUsersController(ui usecase.UsersInteractor) *UsersController {
	return &UsersController{
		usersInteractor: ui,
	}
}

type GetRecommendedUsersResponse struct {
	Users []user
}

func (uc *UsersController) GetRecommendedUsers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	users, err := uc.usersInteractor.GetRecommendedUsers()
	if err != nil {
		return svrErrResp, err
	}

	response := toGetRecommendedUsersResponse(users)

	b, err := json.Marshal(response)
	if err != nil {
		return svrErrResp, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: http.StatusOK,
	}, nil
}

func toGetRecommendedUsersResponse(users []entity.User) GetRecommendedUsersResponse {
	us := make([]user, len(users))
	for i, u := range users {
		us[i] = toUser(u)
	}
	return GetRecommendedUsersResponse{
		Users: us,
	}
}
