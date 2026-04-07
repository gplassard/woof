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

var GetFeatureFlagsEnvironmentCmd = &cobra.Command{
	Use:     "get-feature-flags-environment [environment_id]",
	Aliases: []string{"get-environment"},
	Short:   "Get an environment",
	Long: `Get an environment
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#get-feature-flags-environment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EnvironmentResponse
		var err error

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetFeatureFlagsEnvironment(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-feature-flags-environment")

		fmt.Println(cmdutil.FormatJSON(res, "environments"))
	},
}

func init() {

	Cmd.AddCommand(GetFeatureFlagsEnvironmentCmd)
}
