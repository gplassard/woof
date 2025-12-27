package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteOnCallScheduleCmd = &cobra.Command{
	Use:   "delete_on_call_schedule [schedule_id]",
	Short: "Delete On-Call schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		_, err := api.DeleteOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_on_call_schedule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteOnCallScheduleCmd)
}
