package cmd

import (
	"github.com/spf13/cobra"
)

// Options is command option
type Options struct {
	endpoint string
	user     string
	name     string
	port     string
	profile  string
	region   string
}

var op = &Options{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iam-db-cli",
	Short: "A tool for creating session to rds with iam auth.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Initial is initial command exec from main
func Initial() {
	rootCmd.Flags().StringVarP(&op.endpoint, "endpoint", "e", "", "AWS RDS Endpoint domain name")
	rootCmd.Flags().StringVarP(&op.user, "user", "u", "", "User name for database")
	rootCmd.Flags().StringVarP(&op.name, "name", "n", "", "Database name")
	rootCmd.Flags().StringVarP(&op.port, "port", "P", "3306", "Database port")
	rootCmd.Flags().StringVarP(&op.profile, "profile", "p", "", "AWS IAM Credentials Profile name")
	rootCmd.Flags().StringVarP(&op.region, "region", "r", "ap-northeast-1", "AWS RDS Region name")
	rootCmd.Execute()
}

// GetOptions return option string
func GetOptions() (endpoint string, user string, name string, port string, profile string, region string) {
	return op.endpoint, op.user, op.name, op.port, op.profile, op.region
}
