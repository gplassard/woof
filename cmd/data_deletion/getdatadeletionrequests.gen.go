package data_deletion

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetDataDeletionRequestsCmd = &cobra.Command{
	Use:   "getdatadeletionrequests",
	Short: "Gets a list of data deletion requests",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDataDeletionApi(client.NewAPIClient())
		res, _, err := api.GetDataDeletionRequests(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getdatadeletionrequests: %v", err)
		}

		cmdutil.PrintJSON(res, "data_deletion")
	},
}

func init() {
	Cmd.AddCommand(GetDataDeletionRequestsCmd)
}
