package controller

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pow/skele-api/new_arch/domain/entity"
)

const (
	svrErrMsg      = `{"message": "Internal server error"}`
	notFoundErrMsg = `{"message": "Not Found"}`
)

var svrErrResp = events.APIGatewayProxyResponse{
	Body:       svrErrMsg,
	StatusCode: http.StatusInternalServerError,
}

var notFoundErrResp = events.APIGatewayProxyResponse{
	Body:       notFoundErrMsg,
	StatusCode: http.StatusNotFound,
}

type user struct {
	id int
}

func toUser(u entity.User) user {
	return user{
		id: u.ID,
	}
}
