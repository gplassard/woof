package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentTypeCmd = &cobra.Command{
	Use:     "create-incident-type",
	Aliases: []string{"create-type"},
	Short:   "Create an incident type",
	Long: `Create an incident type
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-type`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTypeResponse
		var err error

		var body datadogV2.IncidentTypeCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncidentType(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-type")

		fmt.Println(cmdutil.FormatJSON(res, "incident_type"))
	},
}

func init() {

	CreateIncidentTypeCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentTypeCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentTypeCmd)
}
