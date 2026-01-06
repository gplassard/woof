package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetLogsArchiveCmd = &cobra.Command{
	Use:     "get-logs-archive [archive_id]",
	Aliases: []string{"get"},
	Short:   "Get an archive",
	Long: `Get an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#get-logs-archive`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsArchive
		var err error

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetLogsArchive(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-logs-archive")

		cmd.Println(cmdutil.FormatJSON(res, "logs_archives"))
	},
}

func init() {

	Cmd.AddCommand(GetLogsArchiveCmd)
}
