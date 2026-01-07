package rum_retention_filters

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var OrderRetentionFiltersCmd = &cobra.Command{
	Use: "order-retention-filters [app_id]",

	Short: "Order RUM retention filters",
	Long: `Order RUM retention filters
Documentation: https://docs.datadoghq.com/api/latest/rum-retention-filters/#order-retention-filters`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumRetentionFiltersOrderResponse
		var err error

		var body datadogV2.RumRetentionFiltersOrderRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.OrderRetentionFilters(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to order-retention-filters")

		fmt.Println(cmdutil.FormatJSON(res, "retention_filters"))
	},
}

func init() {

	OrderRetentionFiltersCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	OrderRetentionFiltersCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(OrderRetentionFiltersCmd)
}
