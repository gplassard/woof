package logs_archives

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AddReadRoleToArchiveCmd = &cobra.Command{
	Use:   "add_read_role_to_archive [archive_id]",
	Short: "Grant role to an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		_, err := api.AddReadRoleToArchive(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToRole{})
		if err != nil {
			log.Fatalf("failed to add_read_role_to_archive: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(AddReadRoleToArchiveCmd)
}
