package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentImpactCmd = &cobra.Command{
	Use:     "create-incident-impact [incident_id]",
	Aliases: []string{"create-impact"},
	Short:   "Create an incident impact",
	Long: `Create an incident impact
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-impact`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentImpactResponse
		var err error

		var body datadogV2.IncidentImpactCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncidentImpact(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-incident-impact")

		fmt.Println(cmdutil.FormatJSON(res, "incident_impact"))
	},
}

func init() {

	CreateIncidentImpactCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentImpactCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentImpactCmd)
}
