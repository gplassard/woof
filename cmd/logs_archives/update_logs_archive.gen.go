package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateLogsArchiveCmd = &cobra.Command{
	Use:     "update-logs-archive [archive_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update an archive",
	Long: `Update an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#update-logs-archive`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsArchive
		var err error

		var body datadogV2.LogsArchiveCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err = api.UpdateLogsArchive(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-logs-archive")

		cmd.Println(cmdutil.FormatJSON(res, "logs_archives"))
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveCmd)
}
