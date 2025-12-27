package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTagConfigurationsCmd = &cobra.Command{
	Use:   "listtagconfigurations",
	Short: "Get a list of metrics",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListTagConfigurations(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listtagconfigurations: %v", err)
		}

		cmdutil.PrintJSON(res, "metrics")
	},
}

func init() {
	Cmd.AddCommand(ListTagConfigurationsCmd)
}
