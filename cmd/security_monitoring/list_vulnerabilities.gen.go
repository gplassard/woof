package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListVulnerabilitiesCmd = &cobra.Command{
	Use: "list-vulnerabilities",

	Short: "List vulnerabilities",
	Long: `List vulnerabilities
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-vulnerabilities`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListVulnerabilitiesResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListVulnerabilities(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-vulnerabilities")

		cmd.Println(cmdutil.FormatJSON(res, "vulnerabilities"))
	},
}

func init() {

	Cmd.AddCommand(ListVulnerabilitiesCmd)
}
