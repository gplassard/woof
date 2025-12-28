package apm_retention_filters

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateApmRetentionFilterCmd = &cobra.Command{
	Use:   "update-apm-retention-filter [filter_id]",
	Aliases: []string{ "update", },
	Short: "Update a retention filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.UpdateApmRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RetentionFilterUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-apm-retention-filter: %v", err)
		}

		cmdutil.PrintJSON(res, "apm_retention_filter")
	},
}

func init() {
	Cmd.AddCommand(UpdateApmRetentionFilterCmd)
}
