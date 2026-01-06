package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetLogsArchiveOrderCmd = &cobra.Command{
	Use:     "get-logs-archive-order",
	Aliases: []string{"get-order"},
	Short:   "Get archive order",
	Long: `Get archive order
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#get-logs-archive-order`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsArchiveOrder
		var err error

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetLogsArchiveOrder(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-logs-archive-order")

		cmd.Println(cmdutil.FormatJSON(res, "archive_order"))
	},
}

func init() {

	Cmd.AddCommand(GetLogsArchiveOrderCmd)
}
