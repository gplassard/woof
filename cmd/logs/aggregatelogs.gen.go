package logs

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AggregateLogsCmd = &cobra.Command{
	Use:   "aggregatelogs",
	Short: "Aggregate events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsApi(client.NewAPIClient())
		res, _, err := api.AggregateLogs(client.NewContext(apiKey, appKey, site), datadogV2.LogsAggregateRequest{})
		if err != nil {
			log.Fatalf("failed to aggregatelogs: %v", err)
		}

		cmdutil.PrintJSON(res, "logs")
	},
}

func init() {
	Cmd.AddCommand(AggregateLogsCmd)
}
