package incidents

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "createincidentnotificationtemplate",
	Short: "Create incident notification template",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), datadogV2.CreateIncidentNotificationTemplateRequest{})
		if err != nil {
			log.Fatalf("failed to createincidentnotificationtemplate: %v", err)
		}

		cmdutil.PrintJSON(res, "incidents")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationTemplateCmd)
}
