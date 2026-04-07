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

var UpdateFeatureFlagCmd = &cobra.Command{
	Use:     "update-feature-flag [feature_flag_id]",
	Aliases: []string{"update"},
	Short:   "Update a feature flag",
	Long: `Update a feature flag
Documentation: https://docs.datadoghq.com/api/latest/feature-flags/#update-feature-flag`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FeatureFlagResponse
		var err error

		var body datadogV2.UpdateFeatureFlagRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFeatureFlagsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateFeatureFlag(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-feature-flag")

		fmt.Println(cmdutil.FormatJSON(res, "feature-flags"))
	},
}

func init() {

	UpdateFeatureFlagCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateFeatureFlagCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateFeatureFlagCmd)
}
