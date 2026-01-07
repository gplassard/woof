package incident_services

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentServiceCmd = &cobra.Command{
	Use:     "update-incident-service [service_id]",
	Aliases: []string{"update"},
	Short:   "Update an existing incident service",
	Long: `Update an existing incident service
Documentation: https://docs.datadoghq.com/api/latest/incident-services/#update-incident-service`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentServiceResponse
		var err error

		var body datadogV2.IncidentServiceUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIncidentService(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-incident-service")

		fmt.Println(cmdutil.FormatJSON(res, "services"))
	},
}

func init() {

	UpdateIncidentServiceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentServiceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIncidentServiceCmd)
}
