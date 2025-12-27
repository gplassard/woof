package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTagConfigurationCmd = &cobra.Command{
	Use:   "createtagconfiguration [metric_name]",
	Short: "Create a tag configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateTagConfiguration(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MetricTagConfigurationCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createtagconfiguration: %v", err)
		}

		cmdutil.PrintJSON(res, "metrics")
	},
}

func init() {
	Cmd.AddCommand(CreateTagConfigurationCmd)
}
