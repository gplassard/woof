package on_call

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteOnCallScheduleCmd = &cobra.Command{
	Use:   "deleteoncallschedule [schedule_id]",
	Short: "Delete On-Call schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		_, err := api.DeleteOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deleteoncallschedule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteOnCallScheduleCmd)
}
