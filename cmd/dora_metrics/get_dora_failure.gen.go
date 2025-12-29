package dora_metrics

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetDORAFailureCmd = &cobra.Command{
	Use:   "get-dora-failure [failure_id]",
	
	Short: "Get a failure event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.GetDORAFailure(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-dora-failure")

		cmd.Println(cmdutil.FormatJSON(res, "dora_metrics"))
	},
}

func init() {
	Cmd.AddCommand(GetDORAFailureCmd)
}
