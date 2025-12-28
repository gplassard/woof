package logs

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListLogsCmd = &cobra.Command{
	Use:   "list-logs",
	Short: "Search logs (POST)",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsApi(client.NewAPIClient())
		res, _, err := api.ListLogs(client.NewContext(apiKey, appKey, site), *datadogV2.NewListLogsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to list-logs: %v", err)
		}

		cmdutil.PrintJSON(res, "log")
	},
}

func init() {
	Cmd.AddCommand(ListLogsCmd)
}
