package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetOnCallScheduleCmd = &cobra.Command{
	Use:   "get-on-call-schedule [schedule_id]",
	Short: "Get On-Call schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.GetOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-on-call-schedule: %v", err)
		}

		cmdutil.PrintJSON(res, "schedules")
	},
}

func init() {
	Cmd.AddCommand(GetOnCallScheduleCmd)
}
