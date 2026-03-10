package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentPostmortemAttachmentCmd = &cobra.Command{
	Use:     "create-incident-postmortem-attachment [incident_id]",
	Aliases: []string{"create-postmortem-attachment"},
	Short:   "Create postmortem attachment",
	Long: `Create postmortem attachment
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-postmortem-attachment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Attachment
		var err error

		var body datadogV2.PostmortemAttachmentRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncidentPostmortemAttachment(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-incident-postmortem-attachment")

		fmt.Println(cmdutil.FormatJSON(res, "incident_attachments"))
	},
}

func init() {

	CreateIncidentPostmortemAttachmentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentPostmortemAttachmentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentPostmortemAttachmentCmd)
}
