package apm_retention_filters

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ReorderApmRetentionFiltersCmd = &cobra.Command{
	Use:   "reorderapmretentionfilters",
	Short: "Re-order retention filters",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		_, err := api.ReorderApmRetentionFilters(client.NewContext(apiKey, appKey, site), datadogV2.ReorderRetentionFiltersRequest{})
		if err != nil {
			log.Fatalf("failed to reorderapmretentionfilters: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ReorderApmRetentionFiltersCmd)
}
