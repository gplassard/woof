package logs

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListLogsGetCmd = &cobra.Command{
	Use:   "list-logs-get",
	Aliases: []string{ "list-get", },
	Short: "Search logs (GET)",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsApi(client.NewAPIClient())
		res, _, err := api.ListLogsGet(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-logs-get")

		cmd.Println(cmdutil.FormatJSON(res, "log"))
	},
}

func init() {
	Cmd.AddCommand(ListLogsGetCmd)
}
