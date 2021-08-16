package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kotaroooo0/lambda-clean-sample/controller/input"
	"github.com/kotaroooo0/lambda-clean-sample/usecase"
)

type LikesController struct {
	likesInteractor usecase.LikesInteractor
}

func NewLikesController(li usecase.LikesInteractor) *LikesController {
	return &LikesController{
		likesInteractor: li,
	}
}

type CreateLikeResponse struct {
	Success bool
}

func (lc *LikesController) CreateLike(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	input, err := input.ParseLikesCreateInput(req)
	if err != nil {
		return svrErrResp, err
	}

	if err := lc.likesInteractor.LikePost(input.PostID, input.UserID); err != nil {
		if err == usecase.ErrNotFound {
			return notFoundErrResp, err
		}
		return svrErrResp, err
		return svrErrResp, err
	}

	response := toCreateLikeResponse(true)

	b, err := json.Marshal(response)
	if err != nil {
		return svrErrResp, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: http.StatusOK,
	}, nil
}

func toCreateLikeResponse(success bool) CreateLikeResponse {
	return CreateLikeResponse{
		Success: success,
	}
}
