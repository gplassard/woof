package reference_tables

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var BatchRowsQueryCmd = &cobra.Command{
	Use: "batch-rows-query",

	Short: "Batch rows query",
	Long: `Batch rows query
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#batch-rows-query`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.BatchRowsQueryResponse
		var err error

		var body datadogV2.BatchRowsQueryRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.BatchRowsQuery(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to batch-rows-query")

		fmt.Println(cmdutil.FormatJSON(res, "reference-tables-batch-rows-query"))
	},
}

func init() {

	BatchRowsQueryCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	BatchRowsQueryCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(BatchRowsQueryCmd)
}
