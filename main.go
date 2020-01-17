package main

import (
	"fmt"
	"iam-db-cli/cmd"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
)

// GenerateDBAuthToken return auth token
func GenerateDBAuthToken(endpoint string, user string, profile string, region string) string {
	awsCredentials := credentials.NewSharedCredentials("", profile)
	authToken, err := rdsutils.BuildAuthToken(
		endpoint,
		region,
		user,
		awsCredentials,
	)
	if err != nil {
		panic(err.Error())
	}
	return authToken
}

func main() {
	cmd.Initial()
	e, u, _, _, p, r := cmd.GetOptions()

	fmt.Println(GenerateDBAuthToken(e, u, p, r))
}
