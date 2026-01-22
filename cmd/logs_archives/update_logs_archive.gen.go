package logs_archives

import (
	"fmt"
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
	Long: `Update an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#update-logs-archive`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsArchive
		var err error

		var body datadogV2.LogsArchiveCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateLogsArchive(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-logs-archive")

		fmt.Println(cmdutil.FormatJSON(res, "logs_archive"))
	},
}

func init() {

	UpdateLogsArchiveCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateLogsArchiveCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateLogsArchiveCmd)
}
