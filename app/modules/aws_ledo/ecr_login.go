package aws_ledo

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"os"
)

func getRegion() string {
	region, exists := os.LookupEnv("AWS_REGION")
	if !exists {
		fmt.Printf("AWS_REGION not set and unable to fetch region")
		os.Exit(1)
	}
	return region
}

func EcrLogin() (*ecr.GetAuthorizationTokenOutput, error) {
	var input *ecr.GetAuthorizationTokenInput
	var token *ecr.GetAuthorizationTokenOutput

	config := &aws.Config{
		Region: aws.String(getRegion()),
	}
	sess, _ := session.NewSession(config)
	_, err := sess.Config.Credentials.Get()
	if err != nil {
		return token, err
	}

	ecrCtx := ecr.New(sess)
	token, err = ecrCtx.GetAuthorizationToken(input)
	return token, nil
}
