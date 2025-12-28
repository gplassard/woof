package aws_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAWSAccountsCmd = &cobra.Command{
	Use:   "list-aws-accounts",
	
	Short: "List all AWS integrations",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListAWSAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-accounts")

		cmdutil.PrintJSON(res, "account")
	},
}

func init() {
	Cmd.AddCommand(ListAWSAccountsCmd)
}
