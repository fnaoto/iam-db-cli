package main

import (
	"fmt"
	"iam-db-cli/cmd"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/go-sql-driver/mysql"
)

// GenerateDBAuthToken return auth token
func GenerateDBAuthToken(endpoint, user, dbname, port, profile, region string) string {
	awsCredentials := credentials.NewSharedCredentials("", profile)
	authToken, err := rdsutils.BuildAuthToken(
		endpoint+":"+port+"/"+dbname,
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
	e, u, n, port, p, r := cmd.GetOptions()

	if e != "" {
		authToken := GenerateDBAuthToken(e, u, n, port, p, r)
		execCmd := fmt.Sprintf("mysql -u %s -h %s -P %s --password='%s' --enable-cleartext-plugin", u, e, port, authToken)
		command := exec.Command("sh", "-c", execCmd)
		command.Stdout = os.Stdout
		command.Stdin = os.Stdin
		command.Stderr = os.Stderr
		command.Run()
	}
}
