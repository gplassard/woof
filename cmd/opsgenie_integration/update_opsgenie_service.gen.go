package opsgenie_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateOpsgenieServiceCmd = &cobra.Command{
	Use: "update-opsgenie-service [integration_service_id] [payload]",

	Short: "Update a single service object",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OpsgenieServiceResponse
		var err error

		var body datadogV2.OpsgenieServiceUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err = api.UpdateOpsgenieService(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-opsgenie-service")

		cmd.Println(cmdutil.FormatJSON(res, "opsgenie-service"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOpsgenieServiceCmd)
}
