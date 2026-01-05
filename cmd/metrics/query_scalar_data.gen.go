package metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var QueryScalarDataCmd = &cobra.Command{
	Use: "query-scalar-data",

	Short: "Query scalar data across multiple products",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.QueryScalarData(client.NewContext(apiKey, appKey, site), datadogV2.ScalarFormulaQueryRequest{})
		cmdutil.HandleError(err, "failed to query-scalar-data")

		cmd.Println(cmdutil.FormatJSON(res, "scalar_response"))
	},
}

func init() {
	Cmd.AddCommand(QueryScalarDataCmd)
}
