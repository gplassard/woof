package test_optimization

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateFlakyTestsCmd = &cobra.Command{
	Use: "update-flaky-tests",

	Short: "Update flaky test states",
	Long: `Update flaky test states
Documentation: https://docs.datadoghq.com/api/latest/test-optimization/#update-flaky-tests`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateFlakyTestsResponse
		var err error

		var body datadogV2.UpdateFlakyTestsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTestOptimizationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateFlakyTests(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-flaky-tests")

		fmt.Println(cmdutil.FormatJSON(res, "update_flaky_test_state_response"))
	},
}

func init() {

	UpdateFlakyTestsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateFlakyTestsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateFlakyTestsCmd)
}
