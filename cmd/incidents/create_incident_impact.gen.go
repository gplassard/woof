package incidents

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentImpactCmd = &cobra.Command{
	Use:   "create-incident-impact [incident_id]",
	Aliases: []string{ "create-impact", },
	Short: "Create an incident impact",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentImpact(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentImpactCreateRequest{})
		cmdutil.HandleError(err, "failed to create-incident-impact")

		cmd.Println(cmdutil.FormatJSON(res, "incident_impacts"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentImpactCmd)
}
