package monitors

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateMonitorConfigPolicyCmd = &cobra.Command{
	Use:   "createmonitorconfigpolicy",
	Short: "Create a monitor configuration policy",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.CreateMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), datadogV2.MonitorConfigPolicyCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createmonitorconfigpolicy: %v", err)
		}

		cmdutil.PrintJSON(res, "monitors")
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorConfigPolicyCmd)
}
