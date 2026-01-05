package logs

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ListLogsCmd = &cobra.Command{
	Use:     "list-logs [payload]",
	Aliases: []string{"list"},
	Short:   "Search logs (POST)",
	Long: `Search logs (POST)
Documentation: https://docs.datadoghq.com/api/latest/logs/#list-logs`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsListResponse
		var err error

		var body datadogV2.ListLogsOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsApi(client.NewAPIClient())
		res, _, err = api.ListLogs(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to list-logs")

		cmd.Println(cmdutil.FormatJSON(res, "log"))
	},
}

func init() {
	Cmd.AddCommand(ListLogsCmd)
}
