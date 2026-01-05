package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateLogsArchiveCmd = &cobra.Command{
	Use:     "create-logs-archive",
	Aliases: []string{"create"},
	Short:   "Create an archive",
	Long: `Create an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#create-logs-archive`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsArchive
		var err error

		var body datadogV2.LogsArchiveCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err = api.CreateLogsArchive(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-logs-archive")

		cmd.Println(cmdutil.FormatJSON(res, "logs_archives"))
	},
}

func init() {

	CreateLogsArchiveCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateLogsArchiveCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateLogsArchiveCmd)
}
