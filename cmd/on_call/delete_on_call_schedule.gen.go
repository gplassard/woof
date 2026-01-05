package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteOnCallScheduleCmd = &cobra.Command{
	Use:     "delete-on-call-schedule [schedule_id]",
	Aliases: []string{"delete-schedule"},
	Short:   "Delete On-Call schedule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		_, err := api.DeleteOnCallSchedule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-on-call-schedule")

	},
}

func init() {
	Cmd.AddCommand(DeleteOnCallScheduleCmd)
}
