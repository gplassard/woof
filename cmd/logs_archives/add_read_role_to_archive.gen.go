package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AddReadRoleToArchiveCmd = &cobra.Command{
	Use: "add-read-role-to-archive [archive_id]",

	Short: "Grant role to an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		_, err := api.AddReadRoleToArchive(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToRole{})
		cmdutil.HandleError(err, "failed to add-read-role-to-archive")

	},
}

func init() {
	Cmd.AddCommand(AddReadRoleToArchiveCmd)
}
