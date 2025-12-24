package aws_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAWSAccountsCmd = &cobra.Command{
	Use:   "listawsaccounts",
	Short: "List all AWS integrations",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListAWSAccounts(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listawsaccounts: %v", err)
		}

		cmdutil.PrintJSON(res, "aws_integration")
	},
}

func init() {
	Cmd.AddCommand(ListAWSAccountsCmd)
}
