package dora_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDORAFailuresCmd = &cobra.Command{
	Use: "list-dora-failures",

	Short: "Get a list of failure events",
	Long: `Get a list of failure events
Documentation: https://docs.datadoghq.com/api/latest/dora-metrics/#list-dora-failures`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORAListResponse
		var err error

		var body datadogV2.DORAListFailuresRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDORAFailures(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to list-dora-failures")

		cmd.Println(cmdutil.FormatJSON(res, "dora_metrics"))
	},
}

func init() {

	ListDORAFailuresCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ListDORAFailuresCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ListDORAFailuresCmd)
}
