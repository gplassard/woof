package logs

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLogsGetCmd = &cobra.Command{
	Use:     "list-logs-get",
	Aliases: []string{"list-get"},
	Short:   "Search logs (GET)",
	Long: `Search logs (GET)
Documentation: https://docs.datadoghq.com/api/latest/logs/#list-logs-get`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsListResponse
		var err error

		api := datadogV2.NewLogsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListLogsGet(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-logs-get")

		fmt.Println(cmdutil.FormatJSON(res, "logs_get"))
	},
}

func init() {

	Cmd.AddCommand(ListLogsGetCmd)
}
