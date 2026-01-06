package downtimes

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDowntimesCmd = &cobra.Command{
	Use:     "list-downtimes",
	Aliases: []string{"list"},
	Short:   "Get all downtimes",
	Long: `Get all downtimes
Documentation: https://docs.datadoghq.com/api/latest/downtimes/#list-downtimes`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListDowntimesResponse
		var err error

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDowntimes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-downtimes")

		cmd.Println(cmdutil.FormatJSON(res, "downtime"))
	},
}

func init() {

	Cmd.AddCommand(ListDowntimesCmd)
}
