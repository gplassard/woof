package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentAttachmentsCmd = &cobra.Command{
	Use:     "update-incident-attachments [incident_id]",
	Aliases: []string{"update-attachments"},
	Short:   "Create, update, and delete incident attachments",
	Long: `Create, update, and delete incident attachments
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident-attachments`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentAttachmentUpdateResponse
		var err error

		var body datadogV2.IncidentAttachmentUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIncidentAttachments(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-incident-attachments")

		fmt.Println(cmdutil.FormatJSON(res, "incident_attachment"))
	},
}

func init() {

	UpdateIncidentAttachmentsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentAttachmentsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIncidentAttachmentsCmd)
}
