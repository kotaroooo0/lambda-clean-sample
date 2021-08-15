package main

import (
	"log"
	"os"

	"github.com/kotaroooo0/lambda-clean-sample/controller"
	"github.com/kotaroooo0/lambda-clean-sample/infrastructure/dynamodb"
	"github.com/kotaroooo0/lambda-clean-sample/infrastructure/firebase"
	"github.com/kotaroooo0/lambda-clean-sample/usecase"
)

func main() {
	dynamodbCli := dynamodb.NewClient()
	firebaseCli, err := firebase.NewMessageingClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	por := dynamodb.NewPostsRepository(*dynamodbCli)
	ur := dynamodb.NewUserRepository(*dynamodbCli)
	pur := firebase.NewPushRepository(firebaseCli)
	li := usecase.NewLikesInteractor(ur, por, pur)
	lc := controller.NewLikesController(li)
	lambda.Start(lc.)
}
