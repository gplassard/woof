package synthetics

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetOnDemandConcurrencyCapCmd = &cobra.Command{
	Use:   "get-on-demand-concurrency-cap",
	
	Short: "Get the on-demand concurrency cap",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		res, _, err := api.GetOnDemandConcurrencyCap(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-on-demand-concurrency-cap")

		cmd.Println(cmdutil.FormatJSON(res, "on_demand_concurrency_cap"))
	},
}

func init() {
	Cmd.AddCommand(GetOnDemandConcurrencyCapCmd)
}
