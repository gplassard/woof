package rum_audience_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var QueryAccountsCmd = &cobra.Command{
	Use: "query-accounts [payload]",

	Short: "Query accounts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.QueryAccountRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.QueryAccounts(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-accounts")

		cmd.Println(cmdutil.FormatJSON(res, "query_response"))
	},
}

func init() {
	Cmd.AddCommand(QueryAccountsCmd)
}
