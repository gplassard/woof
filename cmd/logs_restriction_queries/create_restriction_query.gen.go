package logs_restriction_queries

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRestrictionQueryCmd = &cobra.Command{
	Use: "create-restriction-query",

	Short: "Create a restriction query",
	Long: `Create a restriction query
Documentation: https://docs.datadoghq.com/api/latest/logs-restriction-queries/#create-restriction-query`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionQueryWithoutRelationshipsResponse
		var err error

		var body datadogV2.RestrictionQueryCreatePayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateRestrictionQuery(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-restriction-query")

		fmt.Println(cmdutil.FormatJSON(res, "restriction_query"))
	},
}

func init() {

	CreateRestrictionQueryCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateRestrictionQueryCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateRestrictionQueryCmd)
}
