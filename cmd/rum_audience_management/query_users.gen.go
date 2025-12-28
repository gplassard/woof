package rum_audience_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var QueryUsersCmd = &cobra.Command{
	Use:   "query-users",
	
	Short: "Query users",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.QueryUsers(client.NewContext(apiKey, appKey, site), datadogV2.QueryUsersRequest{})
		cmdutil.HandleError(err, "failed to query-users")

		cmdutil.PrintJSON(res, "query_response")
	},
}

func init() {
	Cmd.AddCommand(QueryUsersCmd)
}
