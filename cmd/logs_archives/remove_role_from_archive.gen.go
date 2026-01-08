package logs_archives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RemoveRoleFromArchiveCmd = &cobra.Command{
	Use: "remove-role-from-archive [archive_id]",

	Short: "Revoke role from an archive",
	Long: `Revoke role from an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#remove-role-from-archive`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.RelationshipToRole
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.RemoveRoleFromArchive(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to remove-role-from-archive")

	},
}

func init() {

	RemoveRoleFromArchiveCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	RemoveRoleFromArchiveCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(RemoveRoleFromArchiveCmd)
}
