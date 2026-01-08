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
	Long: `Grant role to an archive
Documentation: https://docs.datadoghq.com/api/latest/logs-archives/#add-read-role-to-archive`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.RelationshipToRole
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.AddReadRoleToArchive(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to add-read-role-to-archive")

	},
}

func init() {

	AddReadRoleToArchiveCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AddReadRoleToArchiveCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AddReadRoleToArchiveCmd)
}
