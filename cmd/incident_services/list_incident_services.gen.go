package incident_services

import (
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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err := api.ListIncidentServices(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-incident-services")

		cmd.Println(cmdutil.FormatJSON(res, "services"))
	},
}

func init() {
	Cmd.AddCommand(ListIncidentServicesCmd)
}
