package logs_restriction_queries

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRestrictionQueryRolesCmd = &cobra.Command{
	Use: "list-restriction-query-roles [restriction_query_id]",

	Short: "List roles for a restriction query",
	Long: `List roles for a restriction query
Documentation: https://docs.datadoghq.com/api/latest/logs-restriction-queries/#list-restriction-query-roles`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionQueryRolesResponse
		var err error

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err = api.ListRestrictionQueryRoles(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-restriction-query-roles")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {
	Cmd.AddCommand(ListRestrictionQueryRolesCmd)
}
