package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentTypeCmd = &cobra.Command{
	Use:   "createincidenttype",
	Short: "Create an incident type",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentType(client.NewContext(apiKey, appKey, site), datadogV2.IncidentTypeCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createincidenttype: %v", err)
		}

		cmdutil.PrintJSON(res, "incidents")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTypeCmd)
}
