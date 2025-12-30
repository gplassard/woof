package logs_archives

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RemoveRoleFromArchiveCmd = &cobra.Command{
	Use: "remove-role-from-archive [archive_id]",

	Short: "Revoke role from an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		_, err := api.RemoveRoleFromArchive(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToRole{})
		cmdutil.HandleError(err, "failed to remove-role-from-archive")

	},
}

func init() {
	Cmd.AddCommand(RemoveRoleFromArchiveCmd)
}
