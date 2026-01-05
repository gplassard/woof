package downtimes

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateDowntimeCmd = &cobra.Command{
	Use:     "update-downtime [downtime_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update a downtime",
	Long: `Update a downtime
Documentation: https://docs.datadoghq.com/api/latest/downtimes/#update-downtime`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DowntimeResponse
		var err error

		var body datadogV2.DowntimeUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err = api.UpdateDowntime(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-downtime")

		cmd.Println(cmdutil.FormatJSON(res, "downtime"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDowntimeCmd)
}
