package feature_flags

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DisableFeatureFlagEnvironmentCmd = &cobra.Command{
	Use:     "disable-feature-flag-environment [feature_flag_id] [environment_id]",
	Aliases: []string{"disable-environment"},
	Short:   "Disable a feature flag in an environment",
	Long: `Disable a feature flag in an environment
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#disable-feature-flag-environment`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DisableFeatureFlagEnvironment(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), uuid.MustParse(args[1]))
		cmdutil.HandleError(err, "failed to disable-feature-flag-environment")

	},
}

func init() {

	Cmd.AddCommand(DisableFeatureFlagEnvironmentCmd)
}
