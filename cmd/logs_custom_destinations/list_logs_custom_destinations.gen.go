package logs_custom_destinations

import (
	"fmt"
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
	Long: `Get all custom destinations
Documentation: https://docs.datadoghq.com/api/latest/logs-custom-destinations/#list-logs-custom-destinations`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomDestinationsResponse
		var err error

		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListLogsCustomDestinations(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-logs-custom-destinations")

		fmt.Println(cmdutil.FormatJSON(res, "logs_custom_destination"))
	},
}

func init() {

	Cmd.AddCommand(ListLogsCustomDestinationsCmd)
}
