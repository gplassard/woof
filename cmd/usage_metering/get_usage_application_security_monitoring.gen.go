package usage_metering

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"time"

	"github.com/spf13/cobra"
)

var GetUsageApplicationSecurityMonitoringCmd = &cobra.Command{
	Use: "get-usage-application-security-monitoring [start_hr]",

	Short: "Get hourly usage for application security",
	Long: `Get hourly usage for application security
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-usage-application-security-monitoring`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UsageApplicationSecurityMonitoringResponse
		var err error

		optionalParams := datadogV2.NewGetUsageApplicationSecurityMonitoringOptionalParameters()

		if cmd.Flags().Changed("end-hr") {
			val, _ := cmd.Flags().GetString("end-hr")
			optionalParams.WithEndHr(val)
		}

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetUsageApplicationSecurityMonitoring(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), *optionalParams)
		cmdutil.HandleError(err, "failed to get-usage-application-security-monitoring")

		cmd.Println(cmdutil.FormatJSON(res, "usage_timeseries"))
	},
}

func init() {

	GetUsageApplicationSecurityMonitoringCmd.Flags().String("end-hr", "", "Datetime in ISO-8601 format, UTC, precise to hour: '[YYYY-MM-DDThh]' for usage ending **before** this hour.")

	Cmd.AddCommand(GetUsageApplicationSecurityMonitoringCmd)
}
