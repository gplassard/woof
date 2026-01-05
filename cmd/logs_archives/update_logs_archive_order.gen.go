package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateLogsArchiveOrderCmd = &cobra.Command{
	Use:     "update-logs-archive-order [payload]",
	Aliases: []string{"update-order"},
	Short:   "Update archive order",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsArchiveOrder
		var err error

		var body datadogV2.LogsArchiveOrder
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err = api.UpdateLogsArchiveOrder(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-logs-archive-order")

		cmd.Println(cmdutil.FormatJSON(res, "archive_order"))
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveOrderCmd)
}
