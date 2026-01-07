package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentAttachmentsCmd = &cobra.Command{
	Use:     "list-incident-attachments [incident_id]",
	Aliases: []string{"list-attachments"},
	Short:   "List incident attachments",
	Long: `List incident attachments
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incident-attachments`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentAttachmentsResponse
		var err error

		optionalParams := datadogV2.NewListIncidentAttachmentsOptionalParameters()

		if cmd.Flags().Changed("filter-attachment-type") {
			val, _ := cmd.Flags().GetString("filter-attachment-type")
			optionalParams.WithFilterAttachmentType(val)
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentAttachments(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to list-incident-attachments")

		cmd.Println(cmdutil.FormatJSON(res, "incident_attachments"))
	},
}

func init() {

	ListIncidentAttachmentsCmd.Flags().String("filter-attachment-type", "", "Filter attachments by type. Supported values are '1' ('postmortem') and '2' ('link').")

	ListIncidentAttachmentsCmd.Flags().String("include", "", "Resource to include in the response. Supported value: 'last_modified_by_user'.")

	Cmd.AddCommand(ListIncidentAttachmentsCmd)
}
