package incidents

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "updateincidentnotificationtemplate [id]",
	Short: "Update incident notification template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), datadogV2.PatchIncidentNotificationTemplateRequest{})
		if err != nil {
			log.Fatalf("failed to updateincidentnotificationtemplate: %v", err)
		}

		cmdutil.PrintJSON(res, "incidents")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentNotificationTemplateCmd)
}
