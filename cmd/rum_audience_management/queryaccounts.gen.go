package rum_audience_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var QueryAccountsCmd = &cobra.Command{
	Use:   "queryaccounts",
	Short: "Query accounts",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.QueryAccounts(client.NewContext(apiKey, appKey, site), datadogV2.QueryAccountRequest{})
		if err != nil {
			log.Fatalf("failed to queryaccounts: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_audience_management")
	},
}

func init() {
	Cmd.AddCommand(QueryAccountsCmd)
}
