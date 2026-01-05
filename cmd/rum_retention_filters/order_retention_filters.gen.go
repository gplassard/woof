package rum_retention_filters

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var OrderRetentionFiltersCmd = &cobra.Command{
	Use: "order-retention-filters [app_id] [payload]",

	Short: "Order RUM retention filters",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.RumRetentionFiltersOrderRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.OrderRetentionFilters(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to order-retention-filters")

		cmd.Println(cmdutil.FormatJSON(res, "retention_filters"))
	},
}

func init() {
	Cmd.AddCommand(OrderRetentionFiltersCmd)
}
