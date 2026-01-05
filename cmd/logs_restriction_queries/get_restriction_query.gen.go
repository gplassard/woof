package logs_restriction_queries

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRestrictionQueryCmd = &cobra.Command{
	Use: "get-restriction-query [restriction_query_id]",

	Short: "Get a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionQueryWithRelationshipsResponse
		var err error

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err = api.GetRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-restriction-query")

		cmd.Println(cmdutil.FormatJSON(res, "logs_restriction_queries"))
	},
}

func init() {
	Cmd.AddCommand(GetRestrictionQueryCmd)
}
