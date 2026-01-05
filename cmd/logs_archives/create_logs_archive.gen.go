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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.CreateLogsArchive(client.NewContext(apiKey, appKey, site), datadogV2.LogsArchiveCreateRequest{})
		cmdutil.HandleError(err, "failed to create-logs-archive")

		cmd.Println(cmdutil.FormatJSON(res, "logs_archives"))
	},
}

func init() {
	Cmd.AddCommand(CreateLogsArchiveCmd)
}
