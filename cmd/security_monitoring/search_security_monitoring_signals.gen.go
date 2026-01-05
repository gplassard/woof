package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SearchSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:     "search-security-monitoring-signals [payload]",
	Aliases: []string{"search-signals"},
	Short:   "Get a list of security signals",
	Long: `Get a list of security signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#search-security-monitoring-signals`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		var body datadogV2.SearchSecurityMonitoringSignalsOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.SearchSecurityMonitoringSignals(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-security-monitoring-signals")

		cmd.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {
	Cmd.AddCommand(SearchSecurityMonitoringSignalsCmd)
}
