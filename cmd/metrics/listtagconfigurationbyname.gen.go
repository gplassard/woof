package metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTagConfigurationByNameCmd = &cobra.Command{
	Use:   "listtagconfigurationbyname [metric_name]",
	Short: "List tag configuration by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListTagConfigurationByName(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to listtagconfigurationbyname: %v", err)
		}

		cmdutil.PrintJSON(res, "metrics")
	},
}

func init() {
	Cmd.AddCommand(ListTagConfigurationByNameCmd)
}
