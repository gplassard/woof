package on_call

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOnCallScheduleCmd = &cobra.Command{
	Use:   "createoncallschedule",
	Short: "Create On-Call schedule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallSchedule(client.NewContext(apiKey, appKey, site), datadogV2.ScheduleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createoncallschedule: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallScheduleCmd)
}
