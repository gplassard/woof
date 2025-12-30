package opsgenie_integration

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateOpsgenieServiceCmd = &cobra.Command{
	Use: "update-opsgenie-service [integration_service_id]",

	Short: "Update a single service object",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateOpsgenieService(client.NewContext(apiKey, appKey, site), args[0], datadogV2.OpsgenieServiceUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-opsgenie-service")

		cmd.Println(cmdutil.FormatJSON(res, "opsgenie-service"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOpsgenieServiceCmd)
}
