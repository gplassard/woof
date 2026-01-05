package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListArchiveReadRolesCmd = &cobra.Command{
	Use: "list-archive-read-roles [archive_id]",

	Short: "List read roles for an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RolesResponse
		var err error

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err = api.ListArchiveReadRoles(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-archive-read-roles")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {
	Cmd.AddCommand(ListArchiveReadRolesCmd)
}
