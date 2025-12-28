package logs_archives

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteLogsArchiveCmd = &cobra.Command{
	Use:   "delete-logs-archive [archive_id]",
	Aliases: []string{ "delete", },
	Short: "Delete an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		_, err := api.DeleteLogsArchive(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-logs-archive: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteLogsArchiveCmd)
}
