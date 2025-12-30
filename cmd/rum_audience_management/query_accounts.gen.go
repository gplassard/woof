package rum_audience_management

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var QueryAccountsCmd = &cobra.Command{
	Use: "query-accounts",

	Short: "Query accounts",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.QueryAccounts(client.NewContext(apiKey, appKey, site), datadogV2.QueryAccountRequest{})
		cmdutil.HandleError(err, "failed to query-accounts")

		cmd.Println(cmdutil.FormatJSON(res, "query_response"))
	},
}

func init() {
	Cmd.AddCommand(QueryAccountsCmd)
}
