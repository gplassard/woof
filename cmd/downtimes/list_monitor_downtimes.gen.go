package downtimes

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var ListMonitorDowntimesCmd = &cobra.Command{
	Use:     "list-monitor-downtimes [monitor_id]",
	Aliases: []string{"list-monitor"},
	Short:   "Get active downtimes for a monitor",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err := api.ListMonitorDowntimes(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to list-monitor-downtimes")

		cmd.Println(cmdutil.FormatJSON(res, "downtime_match"))
	},
}

func init() {
	Cmd.AddCommand(ListMonitorDowntimesCmd)
}
