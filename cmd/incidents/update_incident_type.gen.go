package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentTypeCmd = &cobra.Command{
	Use:     "update-incident-type [incident_type_id]",
	Aliases: []string{"update-type"},
	Short:   "Update an incident type",
	Long: `Update an incident type
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident-type`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTypeResponse
		var err error

		var body datadogV2.IncidentTypePatchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIncidentType(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-incident-type")

		fmt.Println(cmdutil.FormatJSON(res, "incident_types"))
	},
}

func init() {

	UpdateIncidentTypeCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentTypeCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIncidentTypeCmd)
}
