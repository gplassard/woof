package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetScheduleOnCallUserCmd = &cobra.Command{
	Use:   "get_schedule_on_call_user [schedule_id]",
	Short: "Get scheduled on-call user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.GetScheduleOnCallUser(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get_schedule_on_call_user: %v", err)
		}

		cmdutil.PrintJSON(res, "shifts")
	},
}

func init() {
	Cmd.AddCommand(GetScheduleOnCallUserCmd)
}
