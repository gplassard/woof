package logs_archives

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateLogsArchiveOrderCmd = &cobra.Command{
	Use:   "update-logs-archive-order",
	Short: "Update archive order",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.UpdateLogsArchiveOrder(client.NewContext(apiKey, appKey, site), datadogV2.LogsArchiveOrder{})
		if err != nil {
			log.Fatalf("failed to update-logs-archive-order: %v", err)
		}

		cmdutil.PrintJSON(res, "archive_order")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveOrderCmd)
}
