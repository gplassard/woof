package incident_services

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateIncidentServiceCmd = &cobra.Command{
	Use:     "update-incident-service [service_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update an existing incident service",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentServiceResponse
		var err error

		var body datadogV2.IncidentServiceUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err = api.UpdateIncidentService(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-incident-service")

		cmd.Println(cmdutil.FormatJSON(res, "services"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentServiceCmd)
}
