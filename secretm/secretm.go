package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/jorojas/gambituser/awsgo"
	"github.com/jorojas/gambituser/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson

	fmt.Println(" > Asking for Secret " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)

	fmt.Println(" > Secret lecture OK " + secretName)

	return secretData, nil
}
