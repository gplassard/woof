package test_optimization

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTestOptimizationServiceSettingsCmd = &cobra.Command{
	Use:     "get-test-optimization-service-settings",
	Aliases: []string{"get-service-settings"},
	Short:   "Get Test Optimization service settings",
	Long: `Get Test Optimization service settings
Documentation: https://docs.datadoghq.com/api/latest/test-optimization/#get-test-optimization-service-settings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TestOptimizationServiceSettingsResponse
		var err error

		var body datadogV2.TestOptimizationGetServiceSettingsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTestOptimizationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTestOptimizationServiceSettings(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to get-test-optimization-service-settings")

		fmt.Println(cmdutil.FormatJSON(res, "test_optimization_service_settings"))
	},
}

func init() {

	GetTestOptimizationServiceSettingsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	GetTestOptimizationServiceSettingsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(GetTestOptimizationServiceSettingsCmd)
}
