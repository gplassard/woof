package downtimes

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateDowntimeCmd = &cobra.Command{
	Use:     "update-downtime [downtime_id]",
	Aliases: []string{"update"},
	Short:   "Update a downtime",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err := api.UpdateDowntime(client.NewContext(apiKey, appKey, site), args[0], datadogV2.DowntimeUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-downtime")

		cmd.Println(cmdutil.FormatJSON(res, "downtime"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDowntimeCmd)
}
