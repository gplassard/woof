package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOnCallScheduleCmd = &cobra.Command{
	Use:   "create-on-call-schedule",
	Short: "Create On-Call schedule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallSchedule(client.NewContext(apiKey, appKey, site), datadogV2.ScheduleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-on-call-schedule: %v", err)
		}

		cmdutil.PrintJSON(res, "schedules")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallScheduleCmd)
}
