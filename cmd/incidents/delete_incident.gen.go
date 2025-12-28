package incidents

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteIncidentCmd = &cobra.Command{
	Use:   "delete-incident [incident_id]",
	Aliases: []string{ "delete", },
	Short: "Delete an existing incident",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		_, err := api.DeleteIncident(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-incident")

		
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentCmd)
}
