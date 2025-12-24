package rum_audience_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var QueryUsersCmd = &cobra.Command{
	Use:   "queryusers",
	Short: "Query users",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.QueryUsers(client.NewContext(apiKey, appKey, site), datadogV2.QueryUsersRequest{})
		if err != nil {
			log.Fatalf("failed to queryusers: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_audience_management")
	},
}

func init() {
	Cmd.AddCommand(QueryUsersCmd)
}
