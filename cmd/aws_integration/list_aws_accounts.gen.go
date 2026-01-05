package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAWSAccountsCmd = &cobra.Command{
	Use: "list-aws-accounts",

	Short: "List all AWS integrations",
	Long: `List all AWS integrations
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#list-aws-accounts`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSAccountsResponse
		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err = api.ListAWSAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-accounts")

		cmd.Println(cmdutil.FormatJSON(res, "account"))
	},
}

func init() {

	Cmd.AddCommand(ListAWSAccountsCmd)
}
