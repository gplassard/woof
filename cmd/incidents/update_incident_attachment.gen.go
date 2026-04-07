package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentAttachmentCmd = &cobra.Command{
	Use:     "update-incident-attachment [incident_id] [attachment_id]",
	Aliases: []string{"update-attachment"},
	Short:   "Update incident attachment",
	Long: `Update incident attachment
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident-attachment`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Attachment
		var err error

		var body datadogV2.PatchAttachmentRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIncidentAttachment(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-incident-attachment")

		fmt.Println(cmdutil.FormatJSON(res, "incident_attachments"))
	},
}

func init() {

	UpdateIncidentAttachmentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentAttachmentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIncidentAttachmentCmd)
}
