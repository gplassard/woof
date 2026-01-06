package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentCmd = &cobra.Command{
	Use:     "create-incident",
	Aliases: []string{"create"},
	Short:   "Create an incident",
	Long: `Create an incident
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentResponse
		var err error

		var body datadogV2.IncidentCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncident(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident")

		cmd.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {

	CreateIncidentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentCmd)
}
