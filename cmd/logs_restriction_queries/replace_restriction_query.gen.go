package logs_restriction_queries

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReplaceRestrictionQueryCmd = &cobra.Command{
	Use: "replace-restriction-query [restriction_query_id]",

	Short: "Replace a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err := api.ReplaceRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RestrictionQueryUpdatePayload{})
		cmdutil.HandleError(err, "failed to replace-restriction-query")

		cmd.Println(cmdutil.FormatJSON(res, "logs_restriction_queries"))
	},
}

func init() {
	Cmd.AddCommand(ReplaceRestrictionQueryCmd)
}
