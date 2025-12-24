package rum_retention_filters

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var OrderRetentionFiltersCmd = &cobra.Command{
	Use:   "orderretentionfilters [app_id]",
	Short: "Order RUM retention filters",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.OrderRetentionFilters(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RumRetentionFiltersOrderRequest{})
		if err != nil {
			log.Fatalf("failed to orderretentionfilters: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_retention_filters")
	},
}

func init() {
	Cmd.AddCommand(OrderRetentionFiltersCmd)
}
