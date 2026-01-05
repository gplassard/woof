package logs_archives

import (
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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.UpdateLogsArchiveOrder(client.NewContext(apiKey, appKey, site), datadogV2.LogsArchiveOrder{})
		cmdutil.HandleError(err, "failed to update-logs-archive-order")

		cmd.Println(cmdutil.FormatJSON(res, "archive_order"))
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveOrderCmd)
}
