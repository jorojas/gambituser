package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jorojas/gambituser/awsgo"
	"github.com/jorojas/gambituser/bd"
	"github.com/jorojas/gambituser/models"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()

	if !ValidateParameters() {
		fmt.Println("Error validating parameters. 'SecretManager' MUST be considered.")
		err := errors.New("Error validating parameters. 'SecretManager MUST be considered")

		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Email = " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()

	if err != nil {
		fmt.Println("Error reading secret." + err.Error())
		err := errors.New("Error reading secret." + err.Error())

		return event, err
	}

	err = bd.SignUp(data)

	return event, err
}

func ValidateParameters() bool {
	var getParameter bool
	_, getParameter = os.LookupEnv("SecretName")
	return getParameter
}
