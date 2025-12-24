package downtimes

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDowntimeCmd = &cobra.Command{
	Use:   "updatedowntime [downtime_id]",
	Short: "Update a downtime",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err := api.UpdateDowntime(client.NewContext(apiKey, appKey, site), args[0], datadogV2.DowntimeUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updatedowntime: %v", err)
		}

		cmdutil.PrintJSON(res, "downtimes")
	},
}

func init() {
	Cmd.AddCommand(UpdateDowntimeCmd)
}
