package dynamodb

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/kotaroooo0/lambda-clean-sample/entity"
	"github.com/kotaroooo0/lambda-clean-sample/errors"
)

type UsersRepositoryImpl struct {
	db dynamodb.DynamoDB
}

func NewUserRepository(db dynamodb.DynamoDB) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{
		db: db,
	}
}

const usersTableName = "Users"

var usersTableFields = map[string]*string{
	"#ID": aws.String("Id"),
}

func (ur *UsersRepositoryImpl) Get(userID int) (*entity.User, error) {
	input := &dynamodb.GetItemInput{
		TableName:                aws.String(usersTableName),
		ExpressionAttributeNames: usersTableFields,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {N: aws.String(strconv.Itoa(userID))},
		},
		ProjectionExpression: aws.String("Id"),
	}

	user := entity.User{}
	result, err := ur.db.GetItem(input)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	if len(result.Item) == 0 {
		return nil, &errors.AppError{
			Code:    errors.UserNotFoundCode,
			Message: "user item not found",
		}
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, &user); err != nil {
		return nil, errors.Wrap(err)
	}
	return &user, nil
}

func (ur *UsersRepositoryImpl) GetRecommendedUsers(userID int) ([]entity.User, error) {
	recommended := []entity.User{
		{
			ID: 1,
		},
		{
			ID: 5,
		},
		{
			ID: 7,
		},
	}
	return recommended, nil
}
