package dora_metrics

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDORAFailuresCmd = &cobra.Command{
	Use:   "list-dora-failures",
	
	Short: "Get a list of failure events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.ListDORAFailures(client.NewContext(apiKey, appKey, site), datadogV2.DORAListFailuresRequest{})
		cmdutil.HandleError(err, "failed to list-dora-failures")

		cmd.Println(cmdutil.FormatJSON(res, "dora_metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListDORAFailuresCmd)
}
