package incidents

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetIncidentIntegrationCmd = &cobra.Command{
	Use:   "get-incident-integration [incident_id] [integration_metadata_id]",
	Aliases: []string{ "get-integration", },
	Short: "Get incident integration metadata details",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.GetIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-incident-integration")

		cmd.Println(cmdutil.FormatJSON(res, "incident_integrations"))
	},
}

func init() {
	Cmd.AddCommand(GetIncidentIntegrationCmd)
}
