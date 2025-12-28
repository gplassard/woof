package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateTagConfigurationCmd = &cobra.Command{
	Use:   "update-tag-configuration [metric_name]",
	
	Short: "Update a tag configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.UpdateTagConfiguration(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MetricTagConfigurationUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-tag-configuration: %v", err)
		}

		cmdutil.PrintJSON(res, "manage_tags")
	},
}

func init() {
	Cmd.AddCommand(UpdateTagConfigurationCmd)
}
