package downtimes

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var ListMonitorDowntimesCmd = &cobra.Command{
	Use:     "list-monitor-downtimes [monitor_id]",
	Aliases: []string{"list-monitor"},
	Short:   "Get active downtimes for a monitor",
	Long: `Get active downtimes for a monitor
Documentation: https://docs.datadoghq.com/api/latest/downtimes/#list-monitor-downtimes`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorDowntimeMatchResponse
		var err error

		optionalParams := datadogV2.NewListMonitorDowntimesOptionalParameters()

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListMonitorDowntimes(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), *optionalParams)
		cmdutil.HandleError(err, "failed to list-monitor-downtimes")

		fmt.Println(cmdutil.FormatJSON(res, "downtime_match"))
	},
}

func init() {

	ListMonitorDowntimesCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListMonitorDowntimesCmd.Flags().Int64("page-limit", 0, "Maximum number of downtimes in the response.")

	Cmd.AddCommand(ListMonitorDowntimesCmd)
}
