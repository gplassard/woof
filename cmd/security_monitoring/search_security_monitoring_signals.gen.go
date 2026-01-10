package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:     "search-security-monitoring-signals",
	Aliases: []string{"search-signals"},
	Short:   "Get a list of security signals",
	Long: `Get a list of security signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#search-security-monitoring-signals`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		optionalParams := datadogV2.NewSearchSecurityMonitoringSignalsOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchSecurityMonitoringSignals(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to search-security-monitoring-signals")

		fmt.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {

	SearchSecurityMonitoringSignalsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchSecurityMonitoringSignalsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SearchSecurityMonitoringSignalsCmd)
}
