package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InitializeAWS() {
	Ctx = context.TODO()

	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("Error when loading the configurations from .aws/config" + err.Error())
	}

}
