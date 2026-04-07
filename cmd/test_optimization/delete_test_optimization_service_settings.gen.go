package test_optimization

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTestOptimizationServiceSettingsCmd = &cobra.Command{
	Use:     "delete-test-optimization-service-settings",
	Aliases: []string{"delete-service-settings"},
	Short:   "Delete Test Optimization service settings",
	Long: `Delete Test Optimization service settings
Documentation: https://docs.datadoghq.com/api/latest/test-optimization/#delete-test-optimization-service-settings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.TestOptimizationDeleteServiceSettingsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTestOptimizationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteTestOptimizationServiceSettings(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to delete-test-optimization-service-settings")

	},
}

func init() {

	DeleteTestOptimizationServiceSettingsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteTestOptimizationServiceSettingsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteTestOptimizationServiceSettingsCmd)
}
