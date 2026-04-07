package feature_flags

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var EnableFeatureFlagEnvironmentCmd = &cobra.Command{
	Use:     "enable-feature-flag-environment [feature_flag_id] [environment_id]",
	Aliases: []string{"enable-environment"},
	Short:   "Enable a feature flag in an environment",
	Long: `Enable a feature flag in an environment
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#enable-feature-flag-environment`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.EnableFeatureFlagEnvironment(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), uuid.MustParse(args[1]))
		cmdutil.HandleError(err, "failed to enable-feature-flag-environment")

	},
}

func init() {

	Cmd.AddCommand(EnableFeatureFlagEnvironmentCmd)
}
