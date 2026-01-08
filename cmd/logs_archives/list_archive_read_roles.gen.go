package logs_archives

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListArchiveReadRolesCmd = &cobra.Command{
	Use: "list-archive-read-roles [archive_id]",

	Short: "List read roles for an archive",
	Long: `List read roles for an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#list-archive-read-roles`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RolesResponse
		var err error

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListArchiveReadRoles(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-archive-read-roles")

		fmt.Println(cmdutil.FormatJSON(res, "archive_read_role"))
	},
}

func init() {

	Cmd.AddCommand(ListArchiveReadRolesCmd)
}
