package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentTypesCmd = &cobra.Command{
	Use:     "list-incident-types",
	Aliases: []string{"list-types"},
	Short:   "Get a list of incident types",
	Long: `Get a list of incident types
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incident-types`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTypeListResponse
		var err error

		optionalParams := datadogV2.NewListIncidentTypesOptionalParameters()

		if cmd.Flags().Changed("include-deleted") {
			val, _ := cmd.Flags().GetString("include-deleted")
			optionalParams.WithIncludeDeleted(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentTypes(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-incident-types")

		cmd.Println(cmdutil.FormatJSON(res, "incident_types"))
	},
}

func init() {

	ListIncidentTypesCmd.Flags().String("include-deleted", "", "Include deleted incident types in the response.")

	Cmd.AddCommand(ListIncidentTypesCmd)
}
