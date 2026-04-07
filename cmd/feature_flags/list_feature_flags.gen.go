package feature_flags

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFeatureFlagsCmd = &cobra.Command{
	Use:     "list-feature-flags",
	Aliases: []string{"list"},
	Short:   "List feature flags",
	Long: `List feature flags
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#list-feature-flags`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListFeatureFlagsResponse
		var err error

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListFeatureFlags(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-feature-flags")

		fmt.Println(cmdutil.FormatJSON(res, "feature-flags"))
	},
}

func init() {

	Cmd.AddCommand(ListFeatureFlagsCmd)
}
