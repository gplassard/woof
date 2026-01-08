package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFindingsCmd = &cobra.Command{
	Use: "list-findings",

	Short: "List findings",
	Long: `List findings
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-findings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListFindingsResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListFindings(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-findings")

		fmt.Println(cmdutil.FormatJSON(res, "finding"))
	},
}

func init() {

	Cmd.AddCommand(ListFindingsCmd)
}
