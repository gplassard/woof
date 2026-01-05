package synthetics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SetOnDemandConcurrencyCapCmd = &cobra.Command{
	Use: "set-on-demand-concurrency-cap",

	Short: "Save new value for on-demand concurrency cap",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		res, _, err := api.SetOnDemandConcurrencyCap(client.NewContext(apiKey, appKey, site), datadogV2.OnDemandConcurrencyCapAttributes{})
		cmdutil.HandleError(err, "failed to set-on-demand-concurrency-cap")

		cmd.Println(cmdutil.FormatJSON(res, "on_demand_concurrency_cap"))
	},
}

func init() {
	Cmd.AddCommand(SetOnDemandConcurrencyCapCmd)
}
