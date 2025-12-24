package fleet_automation

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetFleetScheduleCmd = &cobra.Command{
	Use:   "getfleetschedule [id]",
	Short: "Get a schedule by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.GetFleetSchedule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getfleetschedule: %v", err)
		}

		cmdutil.PrintJSON(res, "fleet_automation")
	},
}

func init() {
	Cmd.AddCommand(GetFleetScheduleCmd)
}
