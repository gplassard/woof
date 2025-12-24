package metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteTagConfigurationCmd = &cobra.Command{
	Use:   "deletetagconfiguration [metric_name]",
	Short: "Delete a tag configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		_, err := api.DeleteTagConfiguration(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletetagconfiguration: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteTagConfigurationCmd)
}
