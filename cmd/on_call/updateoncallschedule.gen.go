package on_call

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateOnCallScheduleCmd = &cobra.Command{
	Use:   "updateoncallschedule [schedule_id]",
	Short: "Update On-Call schedule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.UpdateOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ScheduleUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateoncallschedule: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call")
	},
}

func init() {
	Cmd.AddCommand(UpdateOnCallScheduleCmd)
}
