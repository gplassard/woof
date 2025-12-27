package rum_audience_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var QueryEventFilteredUsersCmd = &cobra.Command{
	Use:   "query_event_filtered_users",
	Short: "Query event filtered users",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.QueryEventFilteredUsers(client.NewContext(apiKey, appKey, site), datadogV2.QueryEventFilteredUsersRequest{})
		if err != nil {
			log.Fatalf("failed to query_event_filtered_users: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_audience_management")
	},
}

func init() {
	Cmd.AddCommand(QueryEventFilteredUsersCmd)
}
