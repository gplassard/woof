package dora_metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDORAFailureCmd = &cobra.Command{
	Use: "get-dora-failure [failure_id]",

	Short: "Get a failure event",
	Long: `Get a failure event
Documentation: https://docs.datadoghq.com/api/latest/dora-metrics/#get-dora-failure`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORAFetchResponse
		var err error

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDORAFailure(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-dora-failure")

		fmt.Println(cmdutil.FormatJSON(res, "dora_metrics"))
	},
}

func init() {

	Cmd.AddCommand(GetDORAFailureCmd)
}
