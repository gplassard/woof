package rum_audience_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var QueryEventFilteredUsersCmd = &cobra.Command{
	Use:   "query-event-filtered-users",
	
	Short: "Query event filtered users",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.QueryEventFilteredUsers(client.NewContext(apiKey, appKey, site), datadogV2.QueryEventFilteredUsersRequest{})
		cmdutil.HandleError(err, "failed to query-event-filtered-users")

		cmd.Println(cmdutil.FormatJSON(res, "query_response"))
	},
}

func init() {
	Cmd.AddCommand(QueryEventFilteredUsersCmd)
}
