package logs_restriction_queries

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRestrictionQueriesCmd = &cobra.Command{
	Use: "list-restriction-queries",

	Short: "List restriction queries",
	Long: `List restriction queries
Documentation: https://docs.datadoghq.com/api/latest/logs-restriction-queries/#list-restriction-queries`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionQueryListResponse
		var err error

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err = api.ListRestrictionQueries(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-restriction-queries")

		cmd.Println(cmdutil.FormatJSON(res, "logs_restriction_queries"))
	},
}

func init() {

	Cmd.AddCommand(ListRestrictionQueriesCmd)
}
