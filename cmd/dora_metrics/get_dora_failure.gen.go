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

	Short: "Get an incident event",
	Long: `Get an incident event
Documentation: https://docs.datadoghq.com/api/latest/dora-metrics/#get-dora-failure`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORAFailureFetchResponse
		var err error

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDORAFailure(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-dora-failure")

		fmt.Println(cmdutil.FormatJSON(res, "dora_failure"))
	},
}

func init() {

	Cmd.AddCommand(GetDORAFailureCmd)
}
