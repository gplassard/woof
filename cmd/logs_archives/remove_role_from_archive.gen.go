package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var RemoveRoleFromArchiveCmd = &cobra.Command{
	Use: "remove-role-from-archive [archive_id] [payload]",

	Short: "Revoke role from an archive",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.RelationshipToRole
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		_, err = api.RemoveRoleFromArchive(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to remove-role-from-archive")

	},
}

func init() {
	Cmd.AddCommand(RemoveRoleFromArchiveCmd)
}
