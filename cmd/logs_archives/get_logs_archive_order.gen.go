package logs_archives

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetLogsArchiveOrderCmd = &cobra.Command{
	Use:   "get-logs-archive-order",
	Aliases: []string{ "get-order", },
	Short: "Get archive order",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.GetLogsArchiveOrder(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-logs-archive-order: %v", err)
		}

		cmdutil.PrintJSON(res, "archive_order")
	},
}

func init() {
	Cmd.AddCommand(GetLogsArchiveOrderCmd)
}
