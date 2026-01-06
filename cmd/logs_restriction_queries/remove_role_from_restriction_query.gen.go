package logs_restriction_queries

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RemoveRoleFromRestrictionQueryCmd = &cobra.Command{
	Use: "remove-role-from-restriction-query [restriction_query_id]",

	Short: "Revoke role from a restriction query",
	Long: `Revoke role from a restriction query
Documentation: https://docs.datadoghq.com/api/latest/logs-restriction-queries/#remove-role-from-restriction-query`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.RelationshipToRole
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.RemoveRoleFromRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to remove-role-from-restriction-query")

	},
}

func init() {

	RemoveRoleFromRestrictionQueryCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	RemoveRoleFromRestrictionQueryCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(RemoveRoleFromRestrictionQueryCmd)
}
