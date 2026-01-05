package usage_metering

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"time"

	"github.com/spf13/cobra"
)

var GetUsageApplicationSecurityMonitoringCmd = &cobra.Command{
	Use: "get-usage-application-security-monitoring [start_hr]",

	Short: "Get hourly usage for application security",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetUsageApplicationSecurityMonitoring(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }())
		cmdutil.HandleError(err, "failed to get-usage-application-security-monitoring")

		cmd.Println(cmdutil.FormatJSON(res, "usage_metering"))
	},
}

func init() {
	Cmd.AddCommand(GetUsageApplicationSecurityMonitoringCmd)
}
