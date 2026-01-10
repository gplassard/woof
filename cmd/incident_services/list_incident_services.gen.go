package incident_services

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentServicesCmd = &cobra.Command{
	Use:     "list-incident-services",
	Aliases: []string{"list"},
	Short:   "Get a list of all incident services",
	Long: `Get a list of all incident services
Documentation: https://docs.datadoghq.com/api/latest/incident-services/#list-incident-services`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentServicesResponse
		var err error

		optionalParams := datadogV2.NewListIncidentServicesOptionalParameters()

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

		if cmd.Flags().Changed("filter") {
			val, _ := cmd.Flags().GetString("filter")
			optionalParams.WithFilter(val)
		}

		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentServices(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-incident-services")

		fmt.Println(cmdutil.FormatJSON(res, "services"))
	},
}

func init() {

	ListIncidentServicesCmd.Flags().String("include", "", "Specifies which types of related objects should be included in the response.")

	ListIncidentServicesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListIncidentServicesCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListIncidentServicesCmd.Flags().String("filter", "", "A search query that filters services by name.")

	Cmd.AddCommand(ListIncidentServicesCmd)
}
