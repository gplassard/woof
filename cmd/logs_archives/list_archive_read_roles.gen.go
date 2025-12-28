package logs_archives

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListArchiveReadRolesCmd = &cobra.Command{
	Use:   "list-archive-read-roles [archive_id]",
	
	Short: "List read roles for an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.ListArchiveReadRoles(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list-archive-read-roles: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(ListArchiveReadRolesCmd)
}
