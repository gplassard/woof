package rum_retention_filters

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateRetentionFilterCmd = &cobra.Command{
	Use:   "update-retention-filter [app_id] [rf_id]",
	
	Short: "Update a RUM retention filter",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.UpdateRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.RumRetentionFilterUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-retention-filter")

		cmd.Println(cmdutil.FormatJSON(res, "retention_filters"))
	},
}

func init() {
	Cmd.AddCommand(UpdateRetentionFilterCmd)
}
