package downtimes

import (
	"fmt"
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

		optionalParams := datadogV2.NewListDowntimesOptionalParameters()

		if cmd.Flags().Changed("current-only") {
			val, _ := cmd.Flags().GetString("current-only")
			optionalParams.WithCurrentOnly(val)
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDowntimes(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-downtimes")

		fmt.Println(cmdutil.FormatJSON(res, "downtime"))
	},
}

func init() {

	ListDowntimesCmd.Flags().String("current-only", "", "Only return downtimes that are active when the request is made.")

	ListDowntimesCmd.Flags().String("include", "", "Comma-separated list of resource paths for related resources to include in the response. Supported resource paths are 'created_by' and 'monitor'.")

	ListDowntimesCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListDowntimesCmd.Flags().Int64("page-limit", 0, "Maximum number of downtimes in the response.")

	Cmd.AddCommand(ListDowntimesCmd)
}
