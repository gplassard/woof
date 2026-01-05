package downtimes

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CancelDowntimeCmd = &cobra.Command{
	Use:     "cancel-downtime [downtime_id]",
	Aliases: []string{"cancel"},
	Short:   "Cancel a downtime",
	Long: `Cancel a downtime
Documentation: https://docs.datadoghq.com/api/latest/downtimes/#cancel-downtime`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		_, err = api.CancelDowntime(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to cancel-downtime")

	},
}

func init() {
	Cmd.AddCommand(CancelDowntimeCmd)
}
