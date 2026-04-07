package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteIncidentAttachmentCmd = &cobra.Command{
	Use:     "delete-incident-attachment [incident_id] [attachment_id]",
	Aliases: []string{"delete-attachment"},
	Short:   "Delete incident attachment",
	Long: `Delete incident attachment
Documentation: https://docs.datadoghq.com/api/latest/incidents/#delete-incident-attachment`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteIncidentAttachment(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-incident-attachment")

	},
}

func init() {

	Cmd.AddCommand(DeleteIncidentAttachmentCmd)
}
