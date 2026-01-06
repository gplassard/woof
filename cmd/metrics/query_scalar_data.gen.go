package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var QueryScalarDataCmd = &cobra.Command{
	Use: "query-scalar-data",

	Short: "Query scalar data across multiple products",
	Long: `Query scalar data across multiple products
Documentation: https://docs.datadoghq.com/api/latest/metrics/#query-scalar-data`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ScalarFormulaQueryResponse
		var err error

		var body datadogV2.ScalarFormulaQueryRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.QueryScalarData(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-scalar-data")

		cmd.Println(cmdutil.FormatJSON(res, "scalar_response"))
	},
}

func init() {

	QueryScalarDataCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	QueryScalarDataCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(QueryScalarDataCmd)
}
