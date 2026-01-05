package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateLogsArchiveCmd = &cobra.Command{
	Use:     "create-logs-archive [payload]",
	Aliases: []string{"create"},
	Short:   "Create an archive",
	Long: `Create an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#create-logs-archive`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsArchive
		var err error

		var body datadogV2.LogsArchiveCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err = api.CreateLogsArchive(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-logs-archive")

		cmd.Println(cmdutil.FormatJSON(res, "logs_archives"))
	},
}

func init() {
	Cmd.AddCommand(CreateLogsArchiveCmd)
}
