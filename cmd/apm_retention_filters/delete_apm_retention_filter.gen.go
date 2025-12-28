package apm_retention_filters

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteApmRetentionFilterCmd = &cobra.Command{
	Use:   "delete-apm-retention-filter [filter_id]",
	Aliases: []string{ "delete", },
	Short: "Delete a retention filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		_, err := api.DeleteApmRetentionFilter(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-apm-retention-filter")

		
	},
}

func init() {
	Cmd.AddCommand(DeleteApmRetentionFilterCmd)
}
