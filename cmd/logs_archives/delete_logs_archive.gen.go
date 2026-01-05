package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteLogsArchiveCmd = &cobra.Command{
	Use:     "delete-logs-archive [archive_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an archive",
	Long: `Delete an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#delete-logs-archive`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		_, err = api.DeleteLogsArchive(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-logs-archive")

	},
}

func init() {

	Cmd.AddCommand(DeleteLogsArchiveCmd)
}
