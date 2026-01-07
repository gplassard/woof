package logs_archives

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateLogsArchiveOrderCmd = &cobra.Command{
	Use:     "update-logs-archive-order",
	Aliases: []string{"update-order"},
	Short:   "Update archive order",
	Long: `Update archive order
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#update-logs-archive-order`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsArchiveOrder
		var err error

		var body datadogV2.LogsArchiveOrder
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateLogsArchiveOrder(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-logs-archive-order")

		fmt.Println(cmdutil.FormatJSON(res, "archive_order"))
	},
}

func init() {

	UpdateLogsArchiveOrderCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateLogsArchiveOrderCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateLogsArchiveOrderCmd)
}
