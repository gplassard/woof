package logs_restriction_queries

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListUserRestrictionQueriesCmd = &cobra.Command{
	Use: "list-user-restriction-queries [user_id]",

	Short: "Get all restriction queries for a given user",
	Long: `Get all restriction queries for a given user
Documentation: https://docs.datadoghq.com/api/latest/logs-restriction-queries/#list-user-restriction-queries`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionQueryListResponse
		var err error

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err = api.ListUserRestrictionQueries(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-user-restriction-queries")

		cmd.Println(cmdutil.FormatJSON(res, "logs_restriction_queries"))
	},
}

func init() {

	Cmd.AddCommand(ListUserRestrictionQueriesCmd)
}
