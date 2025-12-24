package apm_retention_filters

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListApmRetentionFiltersCmd = &cobra.Command{
	Use:   "listapmretentionfilters",
	Short: "List all APM retention filters",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.ListApmRetentionFilters(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listapmretentionfilters: %v", err)
		}

		cmdutil.PrintJSON(res, "apm_retention_filters")
	},
}

func init() {
	Cmd.AddCommand(ListApmRetentionFiltersCmd)
}
