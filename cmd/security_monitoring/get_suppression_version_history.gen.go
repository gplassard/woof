package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSuppressionVersionHistoryCmd = &cobra.Command{
	Use: "get-suppression-version-history [suppression_id]",

	Short: "Get a suppression's version history",
	Long: `Get a suppression's version history
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-suppression-version-history`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetSuppressionVersionHistoryResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSuppressionVersionHistory(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-suppression-version-history")

		fmt.Println(cmdutil.FormatJSON(res, "suppression_version_history"))
	},
}

func init() {

	Cmd.AddCommand(GetSuppressionVersionHistoryCmd)
}
