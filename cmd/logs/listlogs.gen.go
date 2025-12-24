package logs

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListLogsCmd = &cobra.Command{
	Use:   "listlogs",
	Short: "Search logs (POST)",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsApi(client.NewAPIClient())
		res, _, err := api.ListLogs(client.NewContext(apiKey, appKey, site), *datadogV2.NewListLogsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to listlogs: %v", err)
		}

		cmdutil.PrintJSON(res, "logs")
	},
}

func init() {
	Cmd.AddCommand(ListLogsCmd)
}
