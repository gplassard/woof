package feature_flags

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFeatureFlagsEnvironmentsCmd = &cobra.Command{
	Use:     "list-feature-flags-environments",
	Aliases: []string{"list-environments"},
	Short:   "List environments",
	Long: `List environments
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#list-feature-flags-environments`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListEnvironmentsResponse
		var err error

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListFeatureFlagsEnvironments(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-feature-flags-environments")

		fmt.Println(cmdutil.FormatJSON(res, "environments"))
	},
}

func init() {

	Cmd.AddCommand(ListFeatureFlagsEnvironmentsCmd)
}
