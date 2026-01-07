package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIncidentCmd = &cobra.Command{
	Use:     "get-incident [incident_id]",
	Aliases: []string{"get"},
	Short:   "Get the details of an incident",
	Long: `Get the details of an incident
Documentation: https://docs.datadoghq.com/api/latest/incidents/#get-incident`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentResponse
		var err error

		optionalParams := datadogV2.NewGetIncidentOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetIncident(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-incident")

		cmd.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {

	GetIncidentCmd.Flags().String("include", "", "Specifies which types of related objects should be included in the response.")

	Cmd.AddCommand(GetIncidentCmd)
}
