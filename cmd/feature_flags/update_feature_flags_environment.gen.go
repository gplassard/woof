package feature_flags

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var UpdateFeatureFlagsEnvironmentCmd = &cobra.Command{
	Use:     "update-feature-flags-environment [environment_id]",
	Aliases: []string{"update-environment"},
	Short:   "Update an environment",
	Long: `Update an environment
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#update-feature-flags-environment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EnvironmentResponse
		var err error

		var body datadogV2.UpdateEnvironmentRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateFeatureFlagsEnvironment(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-feature-flags-environment")

		fmt.Println(cmdutil.FormatJSON(res, "environments"))
	},
}

func init() {

	UpdateFeatureFlagsEnvironmentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateFeatureFlagsEnvironmentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateFeatureFlagsEnvironmentCmd)
}
