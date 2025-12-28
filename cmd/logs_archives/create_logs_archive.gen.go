package logs_archives

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateLogsArchiveCmd = &cobra.Command{
	Use:   "create-logs-archive",
	Aliases: []string{ "create", },
	Short: "Create an archive",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.CreateLogsArchive(client.NewContext(apiKey, appKey, site), datadogV2.LogsArchiveCreateRequest{})
		cmdutil.HandleError(err, "failed to create-logs-archive")

		cmdutil.PrintJSON(res, "logs_archives")
	},
}

func init() {
	Cmd.AddCommand(CreateLogsArchiveCmd)
}
