package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetFindingCmd = &cobra.Command{
	Use: "get-finding [finding_id]",

	Short: "Get a finding",
	Long: `Get a finding
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-finding`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetFindingResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetFinding(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-finding")

		fmt.Println(cmdutil.FormatJSON(res, "detailed_finding"))
	},
}

func init() {

	Cmd.AddCommand(GetFindingCmd)
}
