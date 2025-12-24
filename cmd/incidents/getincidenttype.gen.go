package incidents

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetIncidentTypeCmd = &cobra.Command{
	Use:   "getincidenttype [incident_type_id]",
	Short: "Get incident type details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.GetIncidentType(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getincidenttype: %v", err)
		}

		cmdutil.PrintJSON(res, "incidents")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentTypeCmd)
}
