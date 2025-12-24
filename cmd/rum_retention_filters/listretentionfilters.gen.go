package rum_retention_filters

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRetentionFiltersCmd = &cobra.Command{
	Use:   "listretentionfilters [app_id]",
	Short: "Get all RUM retention filters",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.ListRetentionFilters(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to listretentionfilters: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_retention_filters")
	},
}

func init() {
	Cmd.AddCommand(ListRetentionFiltersCmd)
}
