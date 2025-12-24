package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSecurityMonitoringHistsignalCmd = &cobra.Command{
	Use:   "getsecuritymonitoringhistsignal [histsignal_id]",
	Short: "Get a hist signal's details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSecurityMonitoringHistsignal(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getsecuritymonitoringhistsignal: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringHistsignalCmd)
}
