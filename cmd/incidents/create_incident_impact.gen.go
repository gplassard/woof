package incidents

import (
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

		optionalParams := datadogV2.NewCreateIncidentImpactOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncidentImpact(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to create-incident-impact")

		cmd.Println(cmdutil.FormatJSON(res, "incident_impacts"))
	},
}

func init() {

	CreateIncidentImpactCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentImpactCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	CreateIncidentImpactCmd.Flags().String("include", "", "Specifies which related resources should be included in the response.")

	Cmd.AddCommand(CreateIncidentImpactCmd)
}
