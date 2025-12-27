package downtimes

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var ListMonitorDowntimesCmd = &cobra.Command{
	Use:   "listmonitordowntimes [monitor_id]",
	Short: "Get active downtimes for a monitor",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err := api.ListMonitorDowntimes(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		if err != nil {
			log.Fatalf("failed to listmonitordowntimes: %v", err)
		}

		cmdutil.PrintJSON(res, "downtimes")
	},
}

func init() {
	Cmd.AddCommand(ListMonitorDowntimesCmd)
}
