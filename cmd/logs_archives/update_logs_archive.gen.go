package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateLogsArchiveCmd = &cobra.Command{
	Use:     "update-logs-archive [archive_id]",
	Aliases: []string{"update"},
	Short:   "Update an archive",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.UpdateLogsArchive(client.NewContext(apiKey, appKey, site), args[0], datadogV2.LogsArchiveCreateRequest{})
		cmdutil.HandleError(err, "failed to update-logs-archive")

		cmd.Println(cmdutil.FormatJSON(res, "logs_archives"))
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveCmd)
}
