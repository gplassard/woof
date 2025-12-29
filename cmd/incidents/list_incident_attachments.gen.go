package incidents

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListIncidentAttachmentsCmd = &cobra.Command{
	Use:   "list-incident-attachments [incident_id]",
	Aliases: []string{ "list-attachments", },
	Short: "List incident attachments",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentAttachments(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-incident-attachments")

		cmd.Println(cmdutil.FormatJSON(res, "incident_attachments"))
	},
}

func init() {
	Cmd.AddCommand(ListIncidentAttachmentsCmd)
}
