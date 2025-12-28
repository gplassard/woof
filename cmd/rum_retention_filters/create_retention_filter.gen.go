package rum_retention_filters

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateRetentionFilterCmd = &cobra.Command{
	Use:   "create-retention-filter [app_id]",
	
	Short: "Create a RUM retention filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.CreateRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RumRetentionFilterCreateRequest{})
		cmdutil.HandleError(err, "failed to create-retention-filter")

		cmdutil.PrintJSON(res, "retention_filters")
	},
}

func init() {
	Cmd.AddCommand(CreateRetentionFilterCmd)
}
