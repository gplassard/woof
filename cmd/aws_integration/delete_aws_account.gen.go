package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteAWSAccountCmd = &cobra.Command{
	Use: "delete-aws-account [aws_account_config_id]",

	Short: "Delete an AWS integration",
	Long: `Delete an AWS integration
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#delete-aws-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		_, err = api.DeleteAWSAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-aws-account")

	},
}

func init() {
	Cmd.AddCommand(DeleteAWSAccountCmd)
}
