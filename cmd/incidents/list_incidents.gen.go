package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentsCmd = &cobra.Command{
	Use:     "list-incidents",
	Aliases: []string{"list"},
	Short:   "Get a list of incidents",
	Long: `Get a list of incidents
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incidents`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentsResponse
		var err error

		optionalParams := datadogV2.NewListIncidentsOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidents(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-incidents")

		cmd.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {

	ListIncidentsCmd.Flags().String("include", "", "Specifies which types of related objects should be included in the response.")

	ListIncidentsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListIncidentsCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	Cmd.AddCommand(ListIncidentsCmd)
}
