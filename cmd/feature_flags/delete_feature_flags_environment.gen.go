package feature_flags

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteFeatureFlagsEnvironmentCmd = &cobra.Command{
	Use:     "delete-feature-flags-environment [environment_id]",
	Aliases: []string{"delete-environment"},
	Short:   "Delete an environment",
	Long: `Delete an environment
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#delete-feature-flags-environment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteFeatureFlagsEnvironment(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to delete-feature-flags-environment")

	},
}

func init() {

	Cmd.AddCommand(DeleteFeatureFlagsEnvironmentCmd)
}
