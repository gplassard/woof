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
	Long: `Save new value for on-demand concurrency cap
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#set-on-demand-concurrency-cap`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OnDemandConcurrencyCapResponse
		var err error

		var body datadogV2.OnDemandConcurrencyCapAttributes
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		res, _, err = api.SetOnDemandConcurrencyCap(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to set-on-demand-concurrency-cap")

		cmd.Println(cmdutil.FormatJSON(res, "on_demand_concurrency_cap"))
	},
}

func init() {

	SetOnDemandConcurrencyCapCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SetOnDemandConcurrencyCapCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SetOnDemandConcurrencyCapCmd)
}
