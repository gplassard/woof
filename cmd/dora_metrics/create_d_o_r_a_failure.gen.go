package dora_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDORAFailureCmd = &cobra.Command{
	Use:   "create_d_o_r_a_failure",
	Short: "Send a failure event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateDORAFailure(client.NewContext(apiKey, appKey, site), datadogV2.DORAFailureRequest{})
		if err != nil {
			log.Fatalf("failed to create_d_o_r_a_failure: %v", err)
		}

		cmdutil.PrintJSON(res, "dora_metrics")
	},
}

func init() {
	Cmd.AddCommand(CreateDORAFailureCmd)
}
