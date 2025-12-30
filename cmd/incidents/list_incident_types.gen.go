package incidents

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentTypesCmd = &cobra.Command{
	Use:     "list-incident-types",
	Aliases: []string{"list-types"},
	Short:   "Get a list of incident types",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentTypes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-incident-types")

		cmd.Println(cmdutil.FormatJSON(res, "incident_types"))
	},
}

func init() {
	Cmd.AddCommand(ListIncidentTypesCmd)
}
