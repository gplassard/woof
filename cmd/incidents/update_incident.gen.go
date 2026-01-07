package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentCmd = &cobra.Command{
	Use:     "update-incident [incident_id]",
	Aliases: []string{"update"},
	Short:   "Update an existing incident",
	Long: `Update an existing incident
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentResponse
		var err error

		optionalParams := datadogV2.NewUpdateIncidentOptionalParameters()

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
		res, _, err = api.UpdateIncident(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to update-incident")

		cmd.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {

	UpdateIncidentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	UpdateIncidentCmd.Flags().String("include", "", "Specifies which types of related objects should be included in the response.")

	Cmd.AddCommand(UpdateIncidentCmd)
}
