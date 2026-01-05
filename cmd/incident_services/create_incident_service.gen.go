package incident_services

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentServiceCmd = &cobra.Command{
	Use:     "create-incident-service",
	Aliases: []string{"create"},
	Short:   "Create a new incident service",
	Long: `Create a new incident service
Documentation: https://docs.datadoghq.com/api/latest/incident-services/#create-incident-service`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentServiceResponse
		var err error

		var body datadogV2.IncidentServiceCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err = api.CreateIncidentService(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-service")

		cmd.Println(cmdutil.FormatJSON(res, "services"))
	},
}

func init() {

	CreateIncidentServiceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentServiceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentServiceCmd)
}
