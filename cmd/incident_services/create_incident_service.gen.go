package incident_services

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateIncidentServiceCmd = &cobra.Command{
	Use:     "create-incident-service [payload]",
	Aliases: []string{"create"},
	Short:   "Create a new incident service",
	Long: `Create a new incident service
Documentation: https://docs.datadoghq.com/api/latest/incident-services/#create-incident-service`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentServiceResponse
		var err error

		var body datadogV2.IncidentServiceCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err = api.CreateIncidentService(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-service")

		cmd.Println(cmdutil.FormatJSON(res, "services"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentServiceCmd)
}
