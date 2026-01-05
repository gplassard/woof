package logs_custom_destinations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteLogsCustomDestinationCmd = &cobra.Command{
	Use:     "delete-logs-custom-destination [custom_destination_id]",
	Aliases: []string{"delete"},
	Short:   "Delete a custom destination",
	Long: `Delete a custom destination
Documentation: https://docs.datadoghq.com/api/latest/logs-custom-destinations/#delete-logs-custom-destination`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		_, err = api.DeleteLogsCustomDestination(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-logs-custom-destination")

	},
}

func init() {
	Cmd.AddCommand(DeleteLogsCustomDestinationCmd)
}
