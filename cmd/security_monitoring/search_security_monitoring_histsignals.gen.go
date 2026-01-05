package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SearchSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:     "search-security-monitoring-histsignals [payload]",
	Aliases: []string{"search-histsignals"},
	Short:   "Search hist signals",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		var body datadogV2.SearchSecurityMonitoringHistsignalsOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.SearchSecurityMonitoringHistsignals(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-security-monitoring-histsignals")

		cmd.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {
	Cmd.AddCommand(SearchSecurityMonitoringHistsignalsCmd)
}
