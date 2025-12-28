package fleet_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var TriggerFleetScheduleCmd = &cobra.Command{
	Use:   "trigger-fleet-schedule [id]",
	
	Short: "Trigger a schedule deployment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.TriggerFleetSchedule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to trigger-fleet-schedule: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment")
	},
}

func init() {
	Cmd.AddCommand(TriggerFleetScheduleCmd)
}
