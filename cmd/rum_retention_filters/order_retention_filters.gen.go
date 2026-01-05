package rum_retention_filters

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var OrderRetentionFiltersCmd = &cobra.Command{
	Use: "order-retention-filters [app_id]",

	Short: "Order RUM retention filters",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.OrderRetentionFilters(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RumRetentionFiltersOrderRequest{})
		cmdutil.HandleError(err, "failed to order-retention-filters")

		cmd.Println(cmdutil.FormatJSON(res, "retention_filters"))
	},
}

func init() {
	Cmd.AddCommand(OrderRetentionFiltersCmd)
}
