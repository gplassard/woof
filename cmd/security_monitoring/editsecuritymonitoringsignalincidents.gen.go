package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var EditSecurityMonitoringSignalIncidentsCmd = &cobra.Command{
	Use:   "editsecuritymonitoringsignalincidents [signal_id]",
	Short: "Change the related incidents of a security signal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.EditSecurityMonitoringSignalIncidents(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringSignalIncidentsUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to editsecuritymonitoringsignalincidents: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalIncidentsCmd)
}
