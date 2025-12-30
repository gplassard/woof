package logs

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLogsCmd = &cobra.Command{
	Use:     "list-logs",
	Aliases: []string{"list"},
	Short:   "Search logs (POST)",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsApi(client.NewAPIClient())
		res, _, err := api.ListLogs(client.NewContext(apiKey, appKey, site), *datadogV2.NewListLogsOptionalParameters())
		cmdutil.HandleError(err, "failed to list-logs")

		cmd.Println(cmdutil.FormatJSON(res, "log"))
	},
}

func init() {
	Cmd.AddCommand(ListLogsCmd)
}
