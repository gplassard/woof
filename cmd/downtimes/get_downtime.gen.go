package downtimes

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDowntimeCmd = &cobra.Command{
	Use:     "get-downtime [downtime_id]",
	Aliases: []string{"get"},
	Short:   "Get a downtime",
	Long: `Get a downtime
Documentation: https://docs.datadoghq.com/api/latest/downtimes/#get-downtime`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DowntimeResponse
		var err error

		optionalParams := datadogV2.NewGetDowntimeOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDowntime(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-downtime")

		cmd.Println(cmdutil.FormatJSON(res, "downtime"))
	},
}

func init() {

	GetDowntimeCmd.Flags().String("include", "", "Comma-separated list of resource paths for related resources to include in the response. Supported resource paths are 'created_by' and 'monitor'.")

	Cmd.AddCommand(GetDowntimeCmd)
}
