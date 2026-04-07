package feature_flags

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateFeatureFlagsEnvironmentCmd = &cobra.Command{
	Use:     "create-feature-flags-environment",
	Aliases: []string{"create-environment"},
	Short:   "Create an environment",
	Long: `Create an environment
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#create-feature-flags-environment`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EnvironmentResponse
		var err error

		var body datadogV2.CreateEnvironmentRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateFeatureFlagsEnvironment(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-feature-flags-environment")

		fmt.Println(cmdutil.FormatJSON(res, "environments"))
	},
}

func init() {

	CreateFeatureFlagsEnvironmentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateFeatureFlagsEnvironmentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateFeatureFlagsEnvironmentCmd)
}
