package incident_services

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentServiceCmd = &cobra.Command{
	Use:   "create-incident-service",
	Aliases: []string{ "create", },
	Short: "Create a new incident service",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentService(client.NewContext(apiKey, appKey, site), datadogV2.IncidentServiceCreateRequest{})
		cmdutil.HandleError(err, "failed to create-incident-service")

		cmd.Println(cmdutil.FormatJSON(res, "services"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentServiceCmd)
}
