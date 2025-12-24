package dora_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDORAFailuresCmd = &cobra.Command{
	Use:   "listdorafailures",
	Short: "Get a list of failure events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.ListDORAFailures(client.NewContext(apiKey, appKey, site), datadogV2.DORAListFailuresRequest{})
		if err != nil {
			log.Fatalf("failed to listdorafailures: %v", err)
		}

		cmdutil.PrintJSON(res, "dora_metrics")
	},
}

func init() {
	Cmd.AddCommand(ListDORAFailuresCmd)
}
