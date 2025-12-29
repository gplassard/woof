package logs

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AggregateLogsCmd = &cobra.Command{
	Use:   "aggregate-logs",
	Aliases: []string{ "aggregate", },
	Short: "Aggregate events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsApi(client.NewAPIClient())
		res, _, err := api.AggregateLogs(client.NewContext(apiKey, appKey, site), datadogV2.LogsAggregateRequest{})
		cmdutil.HandleError(err, "failed to aggregate-logs")

		cmd.Println(cmdutil.FormatJSON(res, "logs"))
	},
}

func init() {
	Cmd.AddCommand(AggregateLogsCmd)
}
