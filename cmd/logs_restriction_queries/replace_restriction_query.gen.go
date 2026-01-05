package logs_restriction_queries

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReplaceRestrictionQueryCmd = &cobra.Command{
	Use: "replace-restriction-query [restriction_query_id]",

	Short: "Replace a restriction query",
	Long: `Replace a restriction query
Documentation: https://docs.datadoghq.com/api/latest/logs-restriction-queries/#replace-restriction-query`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionQueryWithoutRelationshipsResponse
		var err error

		var body datadogV2.RestrictionQueryUpdatePayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err = api.ReplaceRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to replace-restriction-query")

		cmd.Println(cmdutil.FormatJSON(res, "logs_restriction_queries"))
	},
}

func init() {

	ReplaceRestrictionQueryCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ReplaceRestrictionQueryCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ReplaceRestrictionQueryCmd)
}
