package logs_archives

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateLogsArchiveCmd = &cobra.Command{
	Use:   "updatelogsarchive [archive_id]",
	Short: "Update an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.UpdateLogsArchive(client.NewContext(apiKey, appKey, site), args[0], datadogV2.LogsArchiveCreateRequest{})
		if err != nil {
			log.Fatalf("failed to updatelogsarchive: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_archives")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveCmd)
}
