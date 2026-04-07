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

var ArchiveFeatureFlagCmd = &cobra.Command{
	Use:     "archive-feature-flag [feature_flag_id]",
	Aliases: []string{"archive"},
	Short:   "Archive a feature flag",
	Long: `Archive a feature flag
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#archive-feature-flag`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FeatureFlagResponse
		var err error

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ArchiveFeatureFlag(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to archive-feature-flag")

		fmt.Println(cmdutil.FormatJSON(res, "feature-flags"))
	},
}

func init() {

	Cmd.AddCommand(ArchiveFeatureFlagCmd)
}
