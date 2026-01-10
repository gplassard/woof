package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:     "search-security-monitoring-histsignals",
	Aliases: []string{"search-histsignals"},
	Short:   "Search hist signals",
	Long: `Search hist signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#search-security-monitoring-histsignals`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		optionalParams := datadogV2.NewSearchSecurityMonitoringHistsignalsOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchSecurityMonitoringHistsignals(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to search-security-monitoring-histsignals")

		fmt.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {

	SearchSecurityMonitoringHistsignalsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchSecurityMonitoringHistsignalsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SearchSecurityMonitoringHistsignalsCmd)
}
