package incident_services

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIncidentServiceCmd = &cobra.Command{
	Use:     "get-incident-service [service_id]",
	Aliases: []string{"get"},
	Short:   "Get details of an incident service",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err := api.GetIncidentService(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-incident-service")

		cmd.Println(cmdutil.FormatJSON(res, "services"))
	},
}

func init() {
	Cmd.AddCommand(GetIncidentServiceCmd)
}
