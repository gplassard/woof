package dora_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteDORAFailureCmd = &cobra.Command{
	Use:   "deletedorafailure [failure_id]",
	Short: "Delete a failure event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		_, err := api.DeleteDORAFailure(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletedorafailure: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteDORAFailureCmd)
}
