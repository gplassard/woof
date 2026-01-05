package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLogsArchivesCmd = &cobra.Command{
	Use:     "list-logs-archives",
	Aliases: []string{"list"},
	Short:   "Get all archives",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.ListLogsArchives(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-logs-archives")

		cmd.Println(cmdutil.FormatJSON(res, "logs_archives"))
	},
}

func init() {
	Cmd.AddCommand(ListLogsArchivesCmd)
}
