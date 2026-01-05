package logs_custom_destinations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLogsCustomDestinationsCmd = &cobra.Command{
	Use:     "list-logs-custom-destinations",
	Aliases: []string{"list"},
	Short:   "Get all custom destinations",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err := api.ListLogsCustomDestinations(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-logs-custom-destinations")

		cmd.Println(cmdutil.FormatJSON(res, "custom_destination"))
	},
}

func init() {
	Cmd.AddCommand(ListLogsCustomDestinationsCmd)
}
