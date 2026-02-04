package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentAttachmentsCmd = &cobra.Command{
	Use:     "list-incident-attachments [incident_id]",
	Aliases: []string{"list-attachments"},
	Short:   "Get a list of attachments",
	Long: `Get a list of attachments
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incident-attachments`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentAttachmentsResponse
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentAttachments(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-incident-attachments")

		fmt.Println(cmdutil.FormatJSON(res, "incident_attachment"))
	},
}

func init() {

	Cmd.AddCommand(ListIncidentAttachmentsCmd)
}
