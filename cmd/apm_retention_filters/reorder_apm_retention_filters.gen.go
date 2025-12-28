package apm_retention_filters

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ReorderApmRetentionFiltersCmd = &cobra.Command{
	Use:   "reorder-apm-retention-filters",
	Aliases: []string{ "reorder", },
	Short: "Re-order retention filters",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		_, err := api.ReorderApmRetentionFilters(client.NewContext(apiKey, appKey, site), datadogV2.ReorderRetentionFiltersRequest{})
		cmdutil.HandleError(err, "failed to reorder-apm-retention-filters")

		
	},
}

func init() {
	Cmd.AddCommand(ReorderApmRetentionFiltersCmd)
}
