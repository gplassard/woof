package logs_archives

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var RemoveRoleFromArchiveCmd = &cobra.Command{
	Use:   "remove_role_from_archive [archive_id]",
	Short: "Revoke role from an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		_, err := api.RemoveRoleFromArchive(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToRole{})
		if err != nil {
			log.Fatalf("failed to remove_role_from_archive: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(RemoveRoleFromArchiveCmd)
}
